package main

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/types"
	"strconv"
	"strings"
)

func (a *goAST) checkType() {
	var (
		info = types.Info{
			Types: make(map[ast.Expr]types.TypeAndValue),
			Defs:  make(map[*ast.Ident]types.Object),
			Uses:  make(map[*ast.Ident]types.Object),
		}
		conf = types.Config{
			Importer: importer.ForCompiler(a.fset, "source", nil),
			// DisableUnusedImportCheck: true,
			Error: func(err error) {
				if sl := strings.Split(err.Error(), "could not import"); len(sl) == 2 {
					fmt.Printf("\n Error: could not import%s\n\n", sl[1])
				}
			},
		}
	)
	conf.Check("check", a.fset, []*ast.File{a.node}, &info)

	ast.Inspect(a.node, func(n ast.Node) bool {
		if n != nil {
			switch x := n.(type) {
			case *ast.CaseClause:
				rewrite(x.Body, &info)
			case *ast.BlockStmt:
				rewrite(x.List, &info)
				// case *ast.CallExpr:
				// return false
			}
		}
		return true
	})
}

func rewrite(in []ast.Stmt, info *types.Info) {
	for k, ex := range in {
		if d, ok := ex.(*ast.DeclStmt); ok {
			if s, ok := d.Decl.(*ast.GenDecl); ok && len(s.Specs) == 1 {
				if v, ok := s.Specs[0].(*ast.ValueSpec); ok && len(v.Names) == 1 {
					var escape bool

					switch {
					case strings.HasPrefix(v.Names[0].Name, "esc"):
						escape = true
					case strings.HasPrefix(v.Names[0].Name, "unesc"):
						escape = false
					default:
						continue
					}

					rewriteValueSpec(in, k, info, v, escape)
				}
			}
		}
	}
}

func rewriteValueSpec(in []ast.Stmt, k int, info *types.Info, v *ast.ValueSpec, escape bool) {
	switch vt := info.TypeOf(v.Values[0]).(type) {
	case *types.Basic:
		if flagVars.stdlib {
			rewriteStdlibBasicType(in, k, vt, v.Values[0], escape)
		} else {
			rewriteBasicType(in, k, vt, v.Values[0], escape)
		}
	default:
		if flagVars.stdlib {
			in[k] = stdlibFuncCall(escape, "fmt", "Sprintf", arg(a(`"%v"`), v.Values[0]))
		} else {
			in[k] = &ast.ExprStmt{X: funcCall(lib_name, "WriteAll", arg(v.Values[0], a(strconv.FormatBool(escape)), a("buffer")))}
		}
	}
}

func rewriteStdlibBasicType(in []ast.Stmt, k int, vt *types.Basic, val ast.Expr, escape bool) {
	switch vt.Name() {
	case "string":
		in[k] = stdlibFuncCall(escape, "", "", arg(val))
	case "int", "int8", "int16", "int32":
		in[k] = stdlibFuncCall(escape, "strconv", "FormatInt", arg(funcCall("", "int64", arg(val)), a("10")))
	case "int64":
		in[k] = stdlibFuncCall(escape, "strconv", "FormatInt", arg(val, a("10")))
	case "uint", "uint8", "uint16", "uint32":
		in[k] = stdlibFuncCall(escape, "strconv", "FormatUint", arg(funcCall("", "uint64", arg(val)), a("10")))
	case "uint64":
		in[k] = stdlibFuncCall(escape, "strconv", "FormatUint", arg(val, a("10")))
	case "bool":
		in[k] = stdlibFuncCall(escape, "strconv", "FormatBool", arg(val))
	case "float64":
		in[k] = stdlibFuncCall(escape, "strconv", "FormatFloat", arg(val, a("'f'"), &ast.UnaryExpr{Op: 13, X: a("1")}, a("64")))
	default:
		in[k] = stdlibFuncCall(escape, "fmt", "Sprintf", arg(a(`"%v"`), val))
	}
}

func rewriteBasicType(in []ast.Stmt, k int, vt *types.Basic, val ast.Expr, escape bool) {
	switch vt.Name() {
	case "string":
		if escape {
			in[k] = &ast.ExprStmt{X: funcCall(lib_name, "WriteEscString", arg(val, a("buffer")))}
		} else {
			in[k] = &ast.ExprStmt{X: funcCall("buffer", "WriteString", arg(val))}
		}
	case "int", "int8", "int16", "int32":
		in[k] = &ast.ExprStmt{X: funcCall(lib_name, "WriteInt", arg(funcCall("", "int64", arg(val)), a("buffer")))}
	case "int64":
		in[k] = &ast.ExprStmt{X: funcCall(lib_name, "WriteInt", arg(val, a("buffer")))}
	case "uint", "uint8", "uint16", "uint32":
		in[k] = &ast.ExprStmt{X: funcCall(lib_name, "WriteUint", arg(funcCall("", "uint64", arg(val)), a("buffer")))}
	case "uint64":
		in[k] = &ast.ExprStmt{X: funcCall(lib_name, "WriteUint", arg(val, a("buffer")))}
	case "bool":
		in[k] = &ast.ExprStmt{X: funcCall(lib_name, "WriteBool", arg(val, a("buffer")))}
	case "float64":
		in[k] = &ast.ExprStmt{X: funcCall(lib_name, "WriteFloat", arg(val, a("buffer")))}
	default:
		in[k] = &ast.ExprStmt{X: funcCall(lib_name, "WriteAll", arg(val, a(strconv.FormatBool(escape)), a("buffer")))}
	}
}

func stdlibFuncCall(esc bool, x, sel string, exps []ast.Expr) *ast.ExprStmt {
	arg := exps
	if sel != "" {
		arg = []ast.Expr{funcCall(x, sel, exps)}
	}
	if esc {
		return &ast.ExprStmt{X: funcCall("buffer", "WriteString", []ast.Expr{funcCall("html", "EscapeString", arg)})}
	} else {
		return &ast.ExprStmt{X: funcCall("buffer", "WriteString", arg)}
	}
}

//

func funcCall(packName, funcName string, exps []ast.Expr) *ast.CallExpr {
	if packName == "" {
		return &ast.CallExpr{
			Fun: &ast.Ident{
				Name: funcName,
			},
			Args: exps,
		}
	}
	return &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   &ast.Ident{Name: packName},
			Sel: &ast.Ident{Name: funcName},
		},
		Args: exps,
	}
}

func arg(i ...ast.Expr) []ast.Expr { return i }
func a(i string) *ast.BasicLit     { return &ast.BasicLit{Value: i} }

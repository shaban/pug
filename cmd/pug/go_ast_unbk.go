package main

import (
	"go/ast"
)

func (a *goAST) checkUnresolvedBlock() {
	var dsSl = []*ast.DeclStmt{}
	var ds *ast.DeclStmt

	ast.Inspect(a.node, func(n ast.Node) bool {
		if n != nil {
			a.handleDeclStmt(n, &dsSl, &ds)
			a.handleCallExpr(n, &dsSl, &ds)
		}
		return true
	})

	ast.Inspect(a.node, func(n ast.Node) bool {
		if n != nil {
			a.handleBlockStmt(n, &dsSl)
		}
		return true
	})
}

func (a *goAST) handleDeclStmt(n ast.Node, dsSl *[]*ast.DeclStmt, ds **ast.DeclStmt) {
	if xds, ok := n.(*ast.DeclStmt); ok {
		if gd, ok := xds.Decl.(*ast.GenDecl); ok {
			ast.Inspect(gd, func(n ast.Node) bool {
				if n != nil {
					if x, ok := n.(*ast.Ident); ok {
						if x.Name == "block" {
							*dsSl = append(*dsSl, xds)
							*ds = xds
						}
					}
				}
				return true
			})
		}
	}
}

func (a *goAST) handleCallExpr(n ast.Node, dsSl *[]*ast.DeclStmt, ds **ast.DeclStmt) {
	if xds, ok := n.(*ast.CallExpr); ok {
		if len(xds.Args) == 1 {
			if i, ok := xds.Args[0].(*ast.Ident); ok {
				if i.Name == "block" {
					if len(*dsSl) > 0 && (*dsSl)[len(*dsSl)-1] == *ds {
						*dsSl = (*dsSl)[:len(*dsSl)-1]
					}
				}
			}
		}
	}
}

func (a *goAST) handleBlockStmt(n ast.Node, dsSl *[]*ast.DeclStmt) {
	if x, ok := n.(*ast.BlockStmt); ok {
		for i, v := range x.List {
			if ds, ok := v.(*ast.DeclStmt); ok {
				for _, v := range *dsSl {
					if ds == v {
						x.List = append(x.List[:i], x.List[i+1:]...)
					}
				}
			}
		}
	}
}

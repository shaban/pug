package main

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"

	"golang.org/x/tools/imports"
)

// goAST struct holds information related to the (AST)
type goAST struct {
	// node: holds a pointer to an ast.File.
	// ast.File is a struct from the go/ast package
	// that represents the root node of the AST for a Go source file.
	// It contains information about the package declaration, imports, other top-level elements of the Go code.
	node *ast.File
	// fset: holds a pointer to a token.FileSet.
	// token.FileSet is a struct from the go/token package
	// that stores information about the source files being parsed,
	// including filename, line numbers, and position information.
	// It's used for associating AST nodes with their location in the source code.
	fset *token.FileSet
}

// bytes purpose is to take the Go AST (Abstract Syntax Tree)
// stored in the goAST struct and convert it into a
// formatted byte slice ([]byte) representing the Go source code.
func (a *goAST) bytes(bb *bytes.Buffer) []byte {
	// uses printer.Fprint function to print the AST.
	// bb: A bytes.Buffer is passed as the first argument. The formatted Go code will be written to this buffer.
	// a.fset: The token.FileSet associated with the AST. This is used to get source file information and line numbers for proper formatting.
	// a.node: The ast.File representing the root node of the AST.
	printer.Fprint(bb, a.fset, a.node)
	// After the AST is printed to the buffer, return the contents of the buffer as byte slice ([]byte) with the formatted Go source code.
	return bb.Bytes()
}

// parseGoSrc is responsible for parsing Go source code and
// creating an Abstract Syntax Tree (AST) representation of that code.
// The AST is a tree-like data structure that represents the syntactic structure of the code.
func parseGoSrc(fileName string, GoSrc interface{}) (out goAST, err error) {
	// Creates a new token.FileSet.
	// A token.FileSet is a data structure that stores information about the source files being parsed,
	// including filename, line numbers, and position information.
	// It's used for associating AST nodes with their location in the source code.
	out.fset = token.NewFileSet()
	// This line does the actual parsing of the Go source code.
	// parser.ParseFile: This function from the go/parser package takes the source code and parses it, creating an AST.
	// fset: The token.FileSet created earlier.
	// fileName: The name of the Go source file.
	// src: The Go source code, which can be a string, a byte slice ([]byte), or an io.Reader.
	// parser.ParseComments: This flag tells the parser to include comments in the AST.
	out.node, err = parser.ParseFile(out.fset, fileName, GoSrc, parser.ParseComments)
	return out, err
}

// goImports function is responsible for taking generated Go source code as a byte slice and then:
// It applies standard Go formatting rules
// It automatically adds, removes, or updates import statements for all necessary packages
// unused imports are removed.
// uses golang.org/x/tools/imports package
func goImports(absPath string, src []byte) []byte {
	// calls the Process function from the imports package, passing the following arguments:
	// absPath: The absolute path of the Go file being processed.
	// This is used for resolving import paths correctly.
	// src: The Go source code as a byte slice ([]byte).
	// &imports.Options{ ... }: An Options struct that configures
	// the behavior of the imports.Process function. Options are:
	// TabWidth: 4: Sets the tab width to 4 spaces.
	// TabIndent: true: Uses tabs for indentation.
	// Comments: true: Preserves comments in the code
	// Fragment: true: Treats the input as a code fragment, which might affect how imports are handled.
	fmtOut, err := imports.Process(absPath, src, &imports.Options{TabWidth: 4, TabIndent: true, Comments: true, Fragment: true})
	// If an error occurred, it logs a fatal error message and terminates the program.
	if err != nil {
		log.Fatalln("goImports(): ", err)
	}
	// If no errors occurred, the function returns the formatted Go code as a byte slice ([]byte).
	return fmtOut
}

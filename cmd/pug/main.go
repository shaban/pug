package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"

	"github.com/shaban/pug"
	"golang.org/x/tools/imports"
)

var (
	dict     = map[string]string{}
	lib_name = ""
	outdir   string
	basedir  string
	pkg_name string
	prepend  string
	stdlib   bool
	stdbuf   bool
	writer   bool
	inline   bool
	// format   bool
	nsFiles = map[string]bool{}
)

func use() {
	fmt.Printf("Usage: %s [OPTION]... [FILE]... \n", os.Args[0])
	flag.PrintDefaults()
}
func init() {
	flag.StringVar(&outdir, "d", "./", `directory for generated .go files`)
	flag.StringVar(&basedir, "basedir", "./", `base directory for templates`)
	flag.StringVar(&pkg_name, "pkg", "pug", `package name for generated files`)
	flag.StringVar(&prepend, "prepend", "", `prepend to generated files e.g for build tags`)
	// flag.BoolVar(&format, "fmt", false, `HTML pretty print output for generated functions`)
	flag.BoolVar(&inline, "inline", false, `inline HTML in generated functions`)
	flag.BoolVar(&stdlib, "stdlib", false, `use stdlib functions`)
	flag.BoolVar(&stdbuf, "stdbuf", false, `use bytes.Buffer  [default bytebufferpool.ByteBuffer]`)
	flag.BoolVar(&writer, "writer", false, `use io.Writer for output`)
}

// goAST struct holds information related to the (AST)
type goAST struct {
	// node: holds a pointer to an ast.File.
	// ast.File is a struct from the go/ast package
	// that represents the root node of the AST for a Go source file.
	// It contains information about the package declaration, imports, declarations, and other top-level elements of the Go code.
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

// genFile takes a Pug template file, parses it,
// generates the corresponding Go code, optimizes and formats the code,
// and then writes the Go code to an output file.
// parameters are the path to the template file for processing
// and the output path where we want our generated go file
func genFile(path, outdir string) (err error) {
	log.Printf("\nfile: %q\n", path)

	var (
		// Splits the provided path into the directory portion (dir) and the filename portion (fname)
		dir, fname = filepath.Split(path)
		// Combines the output directory (outdir) and the filename (fname) to create the full output path for the generated Go file.
		outPath   = outdir + "/" + fname
		rx, _     = regexp.Compile("[^a-zA-Z0-9]+")
		constName = rx.ReplaceAllString(fname[:len(fname)-4], "")
	)
	// get the current working directory and store it as wd
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	// This condition checks if the current working directory (wd)
	// is different from the directory of the Pug file (dir) and if dir is not an empty string.
	if wd != dir && dir != "" {
		// If the condition is true, this line changes the current working directory to the directory of the Pug file.
		os.Chdir(dir)
		// defer changing back to previous directory before genFile returns.
		defer os.Chdir(wd)
	}
	// Checks if a file with the same name (fname) has already been processed.
	if _, ok := nsFiles[fname]; ok {
		// if a collision is found, it appends an underscore and a counter value
		// to the filename and the constant name to avoid naming conflicts.
		sfx := "_" + strconv.Itoa(len(nsFiles))
		// Updates the nsFiles map to track the new filename.
		nsFiles[fname+sfx] = true
		outPath += sfx
		constName += sfx
	} else {
		// If no collision is found, it adds the filename to the nsFiles map.
		nsFiles[fname] = true
	}
	// Reads the contents of the Pug template file (fname) into the fl byte slice.
	fl, err := os.ReadFile(fname)
	if err != nil {
		log.Fatalln("cmd/pug: ReadFile(): ", err)
	}
	// Creates a new Pug parser (pug.New(path)) and uses it to parse the Pug template content (fl).
	// The parsed template is stored in the jst variable.
	jst, err := pug.New(path).Parse(fl)
	if err != nil {
		log.Fatalln("cmd/pug: pug.New(path).Parse(): ", err)
	}

	var (
		bb = new(bytes.Buffer)
		// Creates a new layout struct, which holds information about the Go code structure and imports.
		tpl = newLayout(constName)
	)
	// Writes the initial part of the Go code (package declaration, imports, etc.) to the buffer bb.
	tpl.writeBefore(bb)
	before := bb.Len()
	// Writes the Go code generated from the Pug template to the buffer bb.
	jst.WriteIn(bb)
	// When the template is empty (only consists of the initial part) we're done
	if before == bb.Len() {
		fmt.Print("generated: skipped (empty output)  done.\n\n")
		return
	}
	// Writes the closing parts of the Go code to the buffer bb.
	tpl.writeAfter(bb)

	// Parses the generated Go code in the buffer bb using the parseGoSrc function.
	// This creates an AST (Abstract Syntax Tree) representation of the Go code.
	gst, err := parseGoSrc(outPath, bb)
	// when there is an error write a report to __error.go
	// with the error information and the buffer contents
	// and exit the program indicating a parse error
	if err != nil {
		// TODO
		bb.WriteString("\n\nERROR: parseGoSrc(): ")
		bb.WriteString(err.Error())
		os.WriteFile(outPath+"__Error.go", bb.Bytes(), 0644)
		log.Fatalln("cmd/pug: parseGoSrc(): ", err)
	}
	// Collapses consecutive WriteString calls in the generated Go code to optimize string concatenation.
	gst.collapseWriteString(inline, constName)
	// Performs type checking on the generated Go code to catch potential errors.
	gst.checkType()
	// Checks for unresolved blocks in the generated Go code.
	gst.checkUnresolvedBlock()

	bb.Reset()
	// Uses the goImports function to format the Go code and ensure proper imports.
	fmtOut := goImports(outPath, gst.bytes(bb))
	// Writes the formatted Go code (fmtOut) to the output file
	// (outPath + ".go") with the specified permissions (0644).
	err = os.WriteFile(outPath+".go", fmtOut, 0644)
	if err != nil {
		log.Fatalln("cmd/pug: WriteFile(): ", err)
	}
	fmt.Printf("generated: %s.go  done.\n\n", outPath)
	return nil
}

func genDir(dir, outdir string) {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("\nprevent panic by handling failure accessing path %q: %v", dir, err)
		}

		if ext := filepath.Ext(info.Name()); ext == ".pug" {
			genFile(path, outdir)
		}
		return nil
	})
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	flag.Usage = use // Sets the usage message function
	flag.Parse()     // Parses the command-line arguments
	// Checks if any arguments are provided (exits if not).
	if len(flag.Args()) == 0 {
		use()
		return
	}
	// Config configures the Pug template engine using the settings defined in the golang variable
	// This includes settings
	// for tag and code generation, conditional logic, and mixins.
	pug.Config(golang)
	// Checks if the output directory (outdir) exists; creates it if it doesn't.
	if _, err := os.Stat(outdir); os.IsNotExist(err) {
		// it creates it with the specified permissions
		// (0755, which means read, write, and execute permissions for the owner,
		// and read and execute permissions for group and others).
		os.MkdirAll(outdir, 0755)
	}
	// This line gets the absolute path of the output directory.
	// This is useful to ensure that all subsequent file operations use the full, unambiguous path.
	outdir, _ = filepath.Abs(outdir) // Gets the absolute path of the output directory
	// Changes the working directory to basedir if it's provided
	if _, err := os.Stat(basedir); !os.IsNotExist(err) && basedir != "./" {
		os.Chdir(basedir)
	}
	// Loops through each pugPath provided as a command-line argument
	for _, pugPath := range flag.Args() {
		// get the file descriptor of pugPath for further inspection
		stat, err := os.Stat(pugPath)
		if err != nil {
			log.Fatalln(err)
		}
		// get the absolute path of pugPath
		absPath, _ := filepath.Abs(pugPath)
		// Checks if the pugPath is a directory or a file
		if stat.IsDir() {
			// If it's a directory, calls genDir to process all .pug files in the directory
			genDir(absPath, outdir)
		} else {
			// If it's a file, calls genFile to process the individual .pug file.
			genFile(absPath, outdir)
		}
		// here the decision is made if we want a go template or a pug template
		// at the moment i don't see anything that handles the case when stdlib is true
		// maybe i overlooked something
		if !stdlib {
			makePugFile(stdbuf)
		}
	}
}

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
	//format   bool
	ns_files = map[string]bool{}
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
	//flag.BoolVar(&format, "fmt", false, `HTML pretty print output for generated functions`)
	flag.BoolVar(&inline, "inline", false, `inline HTML in generated functions`)
	flag.BoolVar(&stdlib, "stdlib", false, `use stdlib functions`)
	flag.BoolVar(&stdbuf, "stdbuf", false, `use bytes.Buffer  [default bytebufferpool.ByteBuffer]`)
	flag.BoolVar(&writer, "writer", false, `use io.Writer for output`)
}

//

type goAST struct {
	node *ast.File
	fset *token.FileSet
}

func (a *goAST) bytes(bb *bytes.Buffer) []byte {
	printer.Fprint(bb, a.fset, a.node)
	return bb.Bytes()
}

func parseGoSrc(fileName string, GoSrc interface{}) (out goAST, err error) {
	out.fset = token.NewFileSet()
	out.node, err = parser.ParseFile(out.fset, fileName, GoSrc, parser.ParseComments)
	return
}

func goImports(absPath string, src []byte) []byte {
	fmtOut, err := imports.Process(absPath, src, &imports.Options{TabWidth: 4, TabIndent: true, Comments: true, Fragment: true})
	if err != nil {
		log.Fatalln("goImports(): ", err)
	}

	return fmtOut
}

//

func genFile(path, outdir string) {
	log.Printf("\nfile: %q\n", path)

	var (
		dir, fname = filepath.Split(path)
		outPath    = outdir + "/" + fname
		rx, _      = regexp.Compile("[^a-zA-Z0-9]+")
		constName  = rx.ReplaceAllString(fname[:len(fname)-4], "")
	)

	wd, err := os.Getwd()
	if err == nil && wd != dir && dir != "" {
		os.Chdir(dir)
		defer os.Chdir(wd)
	}

	if _, ok := ns_files[fname]; ok {
		sfx := "_" + strconv.Itoa(len(ns_files))
		ns_files[fname+sfx] = true
		outPath += sfx
		constName += sfx
	} else {
		ns_files[fname] = true
	}

	fl, err := os.ReadFile(fname)
	if err != nil {
		log.Fatalln("cmd/pug: ReadFile(): ", err)
	}

	//

	jst, err := pug.New(path).Parse(fl)
	if err != nil {
		log.Fatalln("cmd/pug: pug.New(path).Parse(): ", err)
	}

	var (
		bb  = new(bytes.Buffer)
		tpl = newLayout(constName)
	)
	tpl.writeBefore(bb)
	before := bb.Len()
	jst.WriteIn(bb)
	if before == bb.Len() {
		fmt.Print("generated: skipped (empty output)  done.\n\n")
		return
	}
	tpl.writeAfter(bb)

	//

	gst, err := parseGoSrc(outPath, bb)
	if err != nil {
		// TODO
		bb.WriteString("\n\nERROR: parseGoSrc(): ")
		bb.WriteString(err.Error())
		os.WriteFile(outPath+"__Error.go", bb.Bytes(), 0644)
		log.Fatalln("cmd/pug: parseGoSrc(): ", err)
	}

	gst.collapseWriteString(inline, constName)
	gst.checkType()
	gst.checkUnresolvedBlock()

	bb.Reset()
	fmtOut := goImports(outPath, gst.bytes(bb))

	//

	err = os.WriteFile(outPath+".go", fmtOut, 0644)
	if err != nil {
		log.Fatalln("cmd/pug: WriteFile(): ", err)
	}
	fmt.Printf("generated: %s.go  done.\n\n", outPath)
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
	flag.Parse()     //Parses the command-line arguments
	//Checks if any arguments are provided (exits if not).
	if len(flag.Args()) == 0 {
		use()
		return
	}
	//configures the Pug template engine using the settings defined in the golang variable
	pug.Config(golang)
	// Checks if the output directory (outdir) exists; creates it if it doesn't.
	if _, err := os.Stat(outdir); os.IsNotExist(err) {
		// it creates it with the specified permissions
		// (0755, which means read, write, and execute permissions for the owner,
		// and read and execute permissions for group and others).
		os.MkdirAll(outdir, 0755)
	}
	//This line gets the absolute path of the output directory.
	//This is useful to ensure that all subsequent file operations use the full, unambiguous path.
	outdir, _ = filepath.Abs(outdir) //Gets the absolute path of the output directory
	//Changes the working directory to basedir if it's provided
	if _, err := os.Stat(basedir); !os.IsNotExist(err) && basedir != "./" {
		os.Chdir(basedir)
	}
	//Loops through each pugPath provided as a command-line argument
	for _, pugPath := range flag.Args() {
		//get the dile descriptor of pugpath for further inspection
		stat, err := os.Stat(pugPath)
		if err != nil {
			log.Fatalln(err)
		}
		//get the absolute path of pugPath
		absPath, _ := filepath.Abs(pugPath)
		//Checks if the pugPath is a directory or a file
		if stat.IsDir() {
			//If it's a directory, calls genDir to process all .pug files in the directory
			genDir(absPath, outdir)
		} else {
			//If it's a file, calls genFile to process the individual .pug file.
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

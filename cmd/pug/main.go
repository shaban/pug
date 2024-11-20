package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
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
	format   bool
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
	flag.BoolVar(&format, "fmt", false, `HTML pretty print output for generated functions`)
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

	fl, err := ioutil.ReadFile(fname)
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
		ioutil.WriteFile(outPath+"__Error.go", bb.Bytes(), 0644)
		log.Fatalln("cmd/pug: parseGoSrc(): ", err)
	}

	gst.collapseWriteString(inline, constName)
	gst.checkType()
	gst.checkUnresolvedBlock()

	bb.Reset()
	fmtOut := goImports(outPath, gst.bytes(bb))

	//

	err = ioutil.WriteFile(outPath+".go", fmtOut, 0644)
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

//

func main() {
	flag.Usage = use
	flag.Parse()
	if len(flag.Args()) == 0 {
		use()
		return
	}

	pug.Config(golang)

	if _, err := os.Stat(outdir); os.IsNotExist(err) {
		os.MkdirAll(outdir, 0755)
	}
	outdir, _ = filepath.Abs(outdir)

	if _, err := os.Stat(basedir); !os.IsNotExist(err) && basedir != "./" {
		os.Chdir(basedir)
	}

	for _, pugPath := range flag.Args() {

		stat, err := os.Stat(pugPath)
		if err != nil {
			log.Fatalln(err)
		}

		absPath, _ := filepath.Abs(pugPath)
		if stat.IsDir() {
			genDir(absPath, outdir)
		} else {
			genFile(absPath, outdir)
		}
		if !stdlib {
			makePugFile(stdbuf)
		}
	}
}

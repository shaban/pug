package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/shaban/pug"
)

var (
	dict     = map[string]string{}
	lib_name = ""
	flagVars struct {
		outdir   string
		basedir  string
		pkg_name string
		prepend  string
		stdlib   bool
		stdbuf   bool
		writer   bool
		inline   bool
		// format   bool
	}

	nsFiles = map[string]bool{}
)

func use() {
	fmt.Printf("Usage: %s [OPTION]... [FILE]... \n", os.Args[0])
	flag.PrintDefaults()
}
func init() {
	flag.StringVar(&flagVars.outdir, "d", "./", `directory for generated .go files`)
	flag.StringVar(&flagVars.basedir, "basedir", "./", `base directory for templates`)
	flag.StringVar(&flagVars.pkg_name, "pkg", "pug", `package name for generated files`)
	flag.StringVar(&flagVars.prepend, "prepend", "", `prepend to generated files e.g for build tags`)
	// flag.BoolVar(&format, "fmt", false, `HTML pretty print output for generated functions`)
	flag.BoolVar(&flagVars.inline, "inline", false, `inline HTML in generated functions`)
	flag.BoolVar(&flagVars.stdlib, "stdlib", false, `use stdlib functions`)
	flag.BoolVar(&flagVars.stdbuf, "stdbuf", false, `use bytes.Buffer  [default bytebufferpool.ByteBuffer]`)
	flag.BoolVar(&flagVars.writer, "writer", false, `use io.Writer for output`)
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
	if _, err := os.Stat(flagVars.outdir); os.IsNotExist(err) {
		// it creates it with the specified permissions
		// (0755, which means read, write, and execute permissions for the owner,
		// and read and execute permissions for group and others).
		os.MkdirAll(flagVars.outdir, 0755)
	}
	// This line gets the absolute path of the output directory.
	// This is useful to ensure that all subsequent file operations use the full, unambiguous path.
	flagVars.outdir, _ = filepath.Abs(flagVars.outdir) // Gets the absolute path of the output directory
	// Changes the working directory to basedir if it's provided
	if _, err := os.Stat(flagVars.basedir); !os.IsNotExist(err) && flagVars.basedir != "./" {
		os.Chdir(flagVars.basedir)
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
			genDir(absPath, flagVars.outdir)
		} else {
			// If it's a file, calls genFile to process the individual .pug file.
			err = genFile(absPath, flagVars.outdir)
			if err != nil {
				log.Fatalln(err)
			}
		}
		// here the decision is made if we want a go template or a pug template
		// at the moment i don't see anything that handles the case when stdlib is true
		if !flagVars.stdlib {
			makePugFile(flagVars.stdbuf)
		}
	}
}

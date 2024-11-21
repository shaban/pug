package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"

	"github.com/shaban/pug"
)

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
		return fmt.Errorf("cmd/pug: Get Working Directory: %s", err.Error())
	}
	// This condition checks if the current working directory (wd)
	// is different from the directory of the Pug file (dir) and if dir is not an empty string.
	if wd != dir && dir != "" {
		// If the condition is true, this line changes the current working directory to the directory of the Pug file.
		err = os.Chdir(dir)
		if err != nil {
			return fmt.Errorf("cmd/pug: Change Working Directory: %s", err.Error())
		}
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
		return fmt.Errorf("cmd/pug: ReadFile(): %s", err.Error())
	}
	// Creates a new Pug parser (pug.New(path)) and uses it to parse the Pug template content (fl).
	// The parsed template is stored in the jst variable.
	jst, err := pug.New(path).Parse(fl)
	if err != nil {
		return fmt.Errorf("cmd/pug: pug.New(path).Parse(): %s", err.Error())
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
		return fmt.Errorf("cmd/pug: parseGoSrc(): %s", err.Error())
	}
	// Collapses consecutive WriteString calls in the generated Go code to optimize string concatenation.
	gst.collapseWriteString(flagVars.inline, constName)
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
		return fmt.Errorf("cmd/pug: WriteFile(): %s", err.Error())
	}
	fmt.Printf("generated: %s.go  done.\n\n", outPath)
	return nil
}

func genDir(dir, outdir string) {
	// uses filepath.Walk function to traverse the directory tree rooted at dir(input directory)
	// filepath.Walk takes a callback function as its second argument.
	// This callback function is called for each file or directory within the directory tree.
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		// The callback function first checks if there was an error during the traversal
		if err != nil {
			// If an error occurred, it returns an error message to prevent a panic.
			return fmt.Errorf("\nprevent panic by handling failure accessing path %q: %v", dir, err)
		}
		// checks if the current file has a .pug extension.
		if ext := filepath.Ext(info.Name()); ext == ".pug" {
			// if the file is a Pug template, it calls the genFile function to process the file
			err = genFile(path, outdir)
			if err != nil {
				log.Fatalln(err)
			}
		}
		return nil
	})
	// After the filepath.Walk function completes, checks if any errors occurred during the traversal.
	// If an error occurred, it logs a fatal error message and terminates the program.
	if err != nil {
		log.Fatalln(err)
	}
}

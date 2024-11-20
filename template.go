// pug.go - template engine. Package implements pug-lang templates for generating Go html/template output.
package pug

import (
	"bytes"
	"io"
	"net/http"
)

/*
Parse parses the template definition string to construct a representation of the template for execution.

Trivial usage:

	package main

	import (
		"fmt"
		"html/template"
		"net/http"

		"github.com/shaban/pug"
	)

	func handler(w http.ResponseWriter, r *http.Request) {
		pugTpl, _ := pug.Parse("pug", []byte("doctype 5\n html: body: p Hello #{.Word}!"))
		goTpl, _ := template.New("html").Parse(pugTpl)

		goTpl.Execute(w, struct{ Word string }{"pug"})
	}

	func main() {
		http.HandleFunc("/", handler)
		http.ListenAndServe(":8080", nil)
	}

Output:

	<!DOCTYPE html><html><body><p>Hello pug!</p></body></html>
*/
func Parse(fname string, text []byte) (string, error) {
	outTpl, err := New(fname).Parse(text)
	if err != nil {
		return "", err
	}
	bb := new(bytes.Buffer)
	outTpl.WriteIn(bb)
	return bb.String(), nil
}

// ParseFile parse the pug template file in given filename
func ParseFile(fname string) (string, error) {
	text, err := ReadFunc(fname)
	if err != nil {
		return "", err
	}
	return Parse(fname, text)
}

// ParseWithFileSystem parse in context of a http.FileSystem (supports embedded files)
func ParseWithFileSystem(fname string, text []byte, fs http.FileSystem) (str string, err error) {
	outTpl := New(fname)
	outTpl.fs = fs

	outTpl, err = outTpl.Parse(text)
	if err != nil {
		return "", err
	}

	bb := new(bytes.Buffer)
	outTpl.WriteIn(bb)
	return bb.String(), nil
}

// ParseFileFromFileSystem parse template file in context of a http.FileSystem (supports embedded files)
func ParseFileFromFileSystem(fname string, fs http.FileSystem) (str string, err error) {
	text, err := readFile(fname, fs)
	if err != nil {
		return "", err
	}
	return ParseWithFileSystem(fname, text, fs)
}

func (t *tree) WriteIn(b io.Writer) {
	t.Root.WriteIn(b)
}

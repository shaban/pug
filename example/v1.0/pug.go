package main

import (
	"fmt"
	"io"
	"os"

	"github.com/shaban/pug"
)

func main() {
	f, err := os.Open("template.pug")
	if err != nil {
		fmt.Printf("Open File error: %v", err)
		return
	}
	dat, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("ReadFile error: %v", err)
		return
	}

	tmpl, err := pug.Parse("name_of_tpl", dat)
	if err != nil {
		fmt.Printf("Parse error: %v", err)
		return
	}

	fmt.Printf("\nOutput:\n\n%s", tmpl)
}

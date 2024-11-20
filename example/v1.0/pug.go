package main

import (
	"fmt"
	"io/ioutil"

	"github.com/Joker/hpp"
	"github.com/shaban/pug"
)

func main() {
	dat, err := ioutil.ReadFile("template.pug")
	if err != nil {
		fmt.Printf("ReadFile error: %v", err)
		return
	}

	tmpl, err := pug.Parse("name_of_tpl", dat)
	if err != nil {
		fmt.Printf("Parse error: %v", err)
		return
	}

	fmt.Printf("\nOutput:\n\n%s", hpp.PrPrint(tmpl))
}

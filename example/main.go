package main

//go:generate pug -pkg=main -writer -fmt index.pug

import (
	"net/http"
)

func main() {

	http.HandleFunc("/", func(wr http.ResponseWriter, req *http.Request) {
		Index("pug.go", true, wr)
	})

	http.ListenAndServe(":8080", nil)
}

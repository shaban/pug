// Code generated by "pug.go"; DO NOT EDIT.

package main

import (
	"io"
)

const (
	index__0 = `<!DOCTYPE html><html lang="en"><head><title>`
	index__1 = `</title><script type="text/javascript">			if(question){
				answer(40 + 2)
			}</script></head><body><h1>pug - template engine`
	index__2 = `</h1><div id="container" class="col">`
	index__3 = `<p>				pug/Pug is a terse and simple
				templating language with
				a <strong>focus</strong> on performance 
				and powerful features.</p></div><footer><div class="footer">2019</div></footer></body></html>`
	index__4 = `<div id="cmd">Precompile pug templates to `
	index__5 = ` code.</div>`
	index__6 = `<p>You are amazing</p>`
	index__7 = `<p>Get on it!</p>`
)

func Index(pageTitle string, youAreUsingpug bool, wr io.Writer) {

	go func() {
		buffer := &WriterAsBuffer{wr}

		buffer.WriteString(index__0)
		WriteEscString(pageTitle, buffer)
		buffer.WriteString(index__1)

		{
			var (
				golang = "Go"
			)

			buffer.WriteString(index__4)
			WriteEscString(golang, buffer)
			buffer.WriteString(index__5)
		}

		buffer.WriteString(index__2)

		if youAreUsingpug {
			buffer.WriteString(index__6)

		} else {
			buffer.WriteString(index__7)

		}
		buffer.WriteString(index__3)
	}()
}

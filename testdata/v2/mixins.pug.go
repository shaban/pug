// Code generated by "pug.go"; DO NOT EDIT.

package pug

import (
	"bytes"

	pool "github.com/valyala/bytebufferpool"
)

const (
	mixins__2  = `<!--  TODO for string -->`
	mixins__3  = `<ul><li>foo</li><li>bar</li><li>baz</li></ul>`
	mixins__5  = `<li class="pet">`
	mixins__11 = `<div class="article"><div class="article-wrapper"><h1>`
	mixins__12 = `</h1>`
	mixins__13 = `</div></div>`
	mixins__14 = `<p>No content provided</p>`
	mixins__18 = `<p>This is my</p><p>Amazing article</p>`
	mixins__21 = `" href="`
	mixins__22 = `">`
	mixins__23 = `</a>`
	mixins__32 = `</h1></div></div>`
	mixins__35 = `<ul id="`
)

func pug_mixins(buffer *pool.ByteBuffer) {

	{
		buffer.WriteString(mixins__3)

	}

	{
		buffer.WriteString(mixins__3)

	}

	buffer.WriteString(iteration__0)
	{
		var (
			name = "cat"
		)

		buffer.WriteString(mixins__5)
		WriteEscString(name, buffer)
		buffer.WriteString(code__7)
	}

	{
		var (
			name = "dog"
		)

		buffer.WriteString(mixins__5)
		WriteEscString(name, buffer)
		buffer.WriteString(code__7)
	}

	{
		var (
			name = "pig"
		)

		buffer.WriteString(mixins__5)
		WriteEscString(name, buffer)
		buffer.WriteString(code__7)
	}

	buffer.WriteString(iteration__3)
	{
		var (
			title = "Hello world"
		)
		var block []byte
		buffer.WriteString(mixins__11)
		WriteEscString(title, buffer)
		buffer.WriteString(mixins__12)
		if len(block) > 0 {
			buffer.Write(block)
		} else {
			buffer.WriteString(mixins__14)

		}
		buffer.WriteString(mixins__13)
	}

	{
		var (
			title = "Hello world"
		)
		var block []byte
		{
			buffer := new(bytes.Buffer)
			buffer.WriteString(mixins__18)

			block = buffer.Bytes()
		}

		buffer.WriteString(mixins__11)
		WriteEscString(title, buffer)
		buffer.WriteString(mixins__12)
		if len(block) > 0 {
			buffer.Write(block)
		} else {
			buffer.WriteString(mixins__14)

		}
		buffer.WriteString(mixins__13)
	}

	{
		var (
			href = "/foo"
			name = "foo"
		)

		attributes := struct{ class string }{class: "btn"}
		buffer.WriteString(attributes__13)
		WriteEscString(attributes.class, buffer)
		buffer.WriteString(mixins__21)
		WriteEscString(href, buffer)
		buffer.WriteString(mixins__22)
		WriteEscString(name, buffer)
		buffer.WriteString(mixins__23)
	}

	{
		var (
			href = fn("/foo", "bar", "baz")
			name = "foo"
		)

		attributes := struct{ class string }{class: "btn"}
		buffer.WriteString(attributes__13)
		WriteEscString(attributes.class, buffer)
		buffer.WriteString(mixins__21)
		WriteAll(href, true, buffer)
		buffer.WriteString(mixins__22)
		WriteEscString(name, buffer)
		buffer.WriteString(mixins__23)
	}

	{
		var (
			href = "/foo"
			name = "foo"
		)

		buffer.WriteString(attributes__0)
		WriteEscString(href, buffer)
		buffer.WriteString(mixins__22)
		WriteEscString(name, buffer)
		buffer.WriteString(mixins__23)
	}

	{
		var (
			title = "Default Title"
		)

		buffer.WriteString(mixins__11)
		WriteEscString(title, buffer)
		buffer.WriteString(mixins__32)

	}

	{
		var (
			title = "Hello world"
		)

		buffer.WriteString(mixins__11)
		WriteEscString(title, buffer)
		buffer.WriteString(mixins__32)

	}

	buffer.WriteString(mixins__2)
	{
		var (
			items = []string{"\"string\"", "2", "3.5", "4"}
			id    = fn("my-list")
		)

		buffer.WriteString(mixins__35)
		WriteAll(id, true, buffer)
		buffer.WriteString(mixins__22)
		for _, item := range items {
			buffer.WriteString(code__6)
			WriteEscString(item, buffer)
			buffer.WriteString(code__7)
		}
		buffer.WriteString(iteration__3)
	}

}

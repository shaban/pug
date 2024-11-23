// Code generated by "pug.go"; DO NOT EDIT.

package pug

import (
	"bytes"

	pool "github.com/valyala/bytebufferpool"
)

const (
	mixinwithblock__0 = `<div></div>`
	mixinwithblock__2 = `<p>test</p>`
)

func pug_mixinwithblock(buffer *pool.ByteBuffer) {

	{
		var block []byte
		buffer.WriteString(mixinwithblock__0)

		if len(block) > 0 {
			buffer.Write(block)
		}
	}

	{
		var block []byte
		{
			buffer := new(bytes.Buffer)
			buffer.WriteString(mixinwithblock__2)

			block = buffer.Bytes()
		}

		buffer.WriteString(mixinwithblock__0)

		if len(block) > 0 {
			buffer.Write(block)
		}
	}

}
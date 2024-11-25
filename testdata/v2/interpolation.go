// Code generated by "pug.go"; DO NOT EDIT.

package pug

import (
	"strings"

	pool "github.com/valyala/bytebufferpool"
)

func pug_interpolation(buffer *pool.ByteBuffer) {

	var title = "On Dogs: Man's Best Friend"
	var author = "enlore"
	var theGreat = "<span>escape!</span>"
	buffer.WriteString(`<h1>`)
	WriteEscString(title, buffer)
	buffer.WriteString(`</h1><p>Written with love by `)
	WriteEscString(author, buffer)
	buffer.WriteString(`</p><p>This will be safe: `)
	WriteEscString(theGreat, buffer)
	buffer.WriteString(`</p>`)
	var msg = "not my inside voice"
	buffer.WriteString(`<p>This is `)
	WriteAll(strings.ToUpper(msg), true, buffer)
	buffer.WriteString(`</p><p>No escaping for `)
	WriteEscString("}", buffer)
	buffer.WriteString(`!</p><p>Escaping works with \#{interpolation}</p><p>Interpolation works with `)
	WriteEscString("#{interpolation}", buffer)
	buffer.WriteString(` too!</p>`)

	var riskyBusiness = "<em>Some of the girls are wearing my mother's clothing.</em>"
	buffer.WriteString(`<div class="quote"><p>Joel: `)
	buffer.WriteString(riskyBusiness)
	buffer.WriteString(`</p></div><p>  This is a very long and boring paragraph that spans multiple lines.
  Suddenly there is a <strong>strongly worded phrase</strong> that cannot be
  <em>ignored</em>.</p><p>  And here's an example of an interpolated tag with an attribute:
  <q lang="es">¡Hola Mundo!</q>

</p><p>If I don't write the paragraph with tag interpolation, tags like<strong>strong</strong>and<em>em</em>might produce unexpected results.</p><p>  If I do, whitespace is <strong>respected</strong> and <em>everybody</em> is happy.
</p>`)

}

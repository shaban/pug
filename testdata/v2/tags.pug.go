// Code generated by "pug.go"; DO NOT EDIT.

package pug

import (
	pool "github.com/valyala/bytebufferpool"
)

const (
	tags__0 = `<ul><li>Item A</li><li>Item B</li><li>Item C</li></ul><img/><a><img/></a><foo/><foo bar="baz"/>`
)

func pug_tags(buffer *pool.ByteBuffer) {

	buffer.WriteString(tags__0)

}
// Code generated by "pug.go"; DO NOT EDIT.

package pug

import (
	pool "github.com/valyala/bytebufferpool"
)

const (
	comments__0 = `<!--  just some paragraphs --><p>foo</p><p>bar</p><p>foo</p><p>bar</p><body><!-- 
    Comments for your HTML readers.
    Use as much text as you want.

 --></body><!DOCTYPE html><!--[if IE 8]><html lang="en" class="lt-ie9"><![endif]--><!--[if gt IE 8]><!--><html lang="en"><!--<![endif]--><body><p>Supporting old web browsers is a pain.</p></body></html>`
)

func pug_comments(buffer *pool.ByteBuffer) {

	buffer.WriteString(comments__0)

}
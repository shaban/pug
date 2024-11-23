// Code generated by "pug.go"; DO NOT EDIT.

package pug

import (
	pool "github.com/valyala/bytebufferpool"
)

const (
	attributes__0 = `<a href="`
	attributes__1 = `">Google</a><a class="button" href="google.com">Google</a><a class="button" href="google.com">Google</a>`
	attributes__2 = `<body class="`
	attributes__3 = `"></body><input type="checkbox" name="agreement" checked="checked"/><input data-json="
  {
    &#34;very-long&#34;: &#34;piece of &#34;,
    &#34;data&#34;: true
  }
"/><!--  pug error --><div class="div-class" (click)="play()"></div><div class="div-class" (click)="play()"></div><div class="div-class" '(click)'="play()"></div><a href="/#{url}">Link`
	attributes__5  = `">Link</a>`
	attributes__7  = `">Another link</a>`
	attributes__8  = `<button type="button" class="`
	attributes__9  = `"></button><button type="button" class="`
	attributes__10 = `"></button></a><div escaped="&lt;code&gt;"></div><div unescaped="<code>"></div><input type="checkbox" checked="checked"/><input type="checkbox" checked="checked"/><input type="checkbox"/><input type="checkbox" checked="true"/><!DOCTYPE html><input type="checkbox" checked="checked"/><input type="checkbox" checked="checked"/><input type="checkbox"/><input type="checkbox" checked="`
	attributes__11 = `"/><a style="`
	attributes__12 = `"></a>`
	attributes__13 = `<a class="`
	attributes__14 = `"></a><a class="bang classes [&#39;bing&#39;]"></a>`
	attributes__16 = `" href="/">Home</a><a class="`
	attributes__17 = `" href="/about">About</a><a class="button"></a><div class="content"></div><a id="main-link"></a><div id="content"></div><div id="foo" data-bar="foo"></div>`
	attributes__18 = `<div id="foo" data-bar="foo"></div><zxc class="asd qwe zxc" num="`
	attributes__19 = `"></zxc><zxc num="`
	attributes__20 = `"></zxc>`
)

func pug_attributes(buffer *pool.ByteBuffer) {

	buffer.WriteString(attributes__0)
	WriteEscString(`google.com`+`google.com`, buffer)
	buffer.WriteString(attributes__1)

	var authenticated = true
	buffer.WriteString(attributes__2)
	WriteAll(ternary(authenticated, "authed", "anon"), true, buffer)
	buffer.WriteString(attributes__3)

	var url = "pug-test.html"
	buffer.WriteString(attributes__0)
	WriteEscString("/"+url, buffer)
	buffer.WriteString(attributes__5)

	url = "https://example.com/"
	buffer.WriteString(attributes__0)
	WriteEscString(url, buffer)
	buffer.WriteString(attributes__7)

	var btnType = "info"
	var btnSize = "lg"
	buffer.WriteString(attributes__8)
	WriteEscString("btn btn-"+btnType+" btn-"+btnSize, buffer)
	buffer.WriteString(attributes__9)
	WriteEscString(`btn btn-`+btnType+` btn-`+btnSize+``, buffer)
	buffer.WriteString(attributes__10)
	WriteBool(true && "checked" == "checked", buffer)
	buffer.WriteString(attributes__11)
	WriteAll(map[string]string{"color": "red", "background": "green"}, true, buffer)
	buffer.WriteString(attributes__12)

	var classes = []string{"foo", "bar", "baz"}
	buffer.WriteString(attributes__13)
	WriteAll(classes, true, buffer)
	buffer.WriteString(attributes__14)

	var currentUrl = "/about"
	buffer.WriteString(attributes__13)
	WriteAll(ternary(currentUrl == "/", "active", ""), true, buffer)
	buffer.WriteString(attributes__16)
	WriteAll(ternary(currentUrl == "/about", "active", ""), true, buffer)
	buffer.WriteString(attributes__17)

	var attributes = struct{ class string }{}
	attributes.class = "baz"
	buffer.WriteString(attributes__18)
	WriteInt(int64(1), buffer)
	buffer.WriteString(attributes__19)
	WriteFloat(1.1, buffer)
	buffer.WriteString(attributes__20)

}
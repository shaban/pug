// Code generated by "pug.go"; DO NOT EDIT.

package pug

import (
	pool "github.com/valyala/bytebufferpool"
)

func pug_attributes(buffer *pool.ByteBuffer) {

	buffer.WriteString(`<a href="`)
	WriteEscString(`google.com`+`google.com`, buffer)
	buffer.WriteString(`">Google</a><a class="button" href="google.com">Google</a><a class="button" href="google.com">Google</a>`)

	var authenticated = true
	buffer.WriteString(`<body class="`)
	WriteAll(ternary(authenticated, "authed", "anon"), true, buffer)
	buffer.WriteString(`"></body><input type="checkbox" name="agreement" checked="checked"/><input data-json="
  {
    &#34;very-long&#34;: &#34;piece of &#34;,
    &#34;data&#34;: true
  }
"/><!--  pug error --><div class="div-class" (click)="play()"></div><div class="div-class" (click)="play()"></div><div class="div-class" '(click)'="play()"></div><a href="/#{url}">Link`)

	var url = "pug-test.html"
	buffer.WriteString(`<a href="`)
	WriteEscString("/"+url, buffer)
	buffer.WriteString(`">Link</a>`)

	url = "https://example.com/"
	buffer.WriteString(`<a href="`)
	WriteEscString(url, buffer)
	buffer.WriteString(`">Another link</a>`)

	var btnType = "info"
	var btnSize = "lg"
	buffer.WriteString(`<button type="button" class="`)
	WriteEscString("btn btn-"+btnType+" btn-"+btnSize, buffer)
	buffer.WriteString(`"></button><button type="button" class="`)
	WriteEscString(`btn btn-`+btnType+` btn-`+btnSize+``, buffer)
	buffer.WriteString(`"></button></a><div escaped="&lt;code&gt;"></div><div unescaped="<code>"></div><input type="checkbox" checked="checked"/><input type="checkbox" checked="checked"/><input type="checkbox"/><input type="checkbox" checked="true"/><!DOCTYPE html><input type="checkbox" checked="checked"/><input type="checkbox" checked="checked"/><input type="checkbox"/><input type="checkbox" checked="`)
	WriteBool(true && "checked" == "checked", buffer)
	buffer.WriteString(`"/><a style="`)
	WriteAll(map[string]string{"color": "red", "background": "green"}, true, buffer)
	buffer.WriteString(`"></a>`)

	var classes = []string{"foo", "bar", "baz"}
	buffer.WriteString(`<a class="`)
	WriteAll(classes, true, buffer)
	buffer.WriteString(`"></a><a class="bang classes [&#39;bing&#39;]"></a>`)

	var currentUrl = "/about"
	buffer.WriteString(`<a class="`)
	WriteAll(ternary(currentUrl == "/", "active", ""), true, buffer)
	buffer.WriteString(`" href="/">Home</a><a class="`)
	WriteAll(ternary(currentUrl == "/about", "active", ""), true, buffer)
	buffer.WriteString(`" href="/about">About</a><a class="button"></a><div class="content"></div><a id="main-link"></a><div id="content"></div><div id="foo" data-bar="foo"></div>`)

	var attributes = struct{ class string }{}
	attributes.class = "baz"
	buffer.WriteString(`<div id="foo" data-bar="foo"></div><zxc class="asd qwe zxc" num="`)
	WriteInt(int64(1), buffer)
	buffer.WriteString(`"></zxc><zxc num="`)
	WriteFloat(1.1, buffer)
	buffer.WriteString(`"></zxc>`)

}

// Code generated by "pug.go"; DO NOT EDIT.

package pug

import (
	"strconv"

	pool "github.com/valyala/bytebufferpool"
)

func pug_iteration(buffer *pool.ByteBuffer) {

	buffer.WriteString(`<ul>`)
	for _, val := range []int{1, 2, 3, 4, 5} {
		buffer.WriteString(`<li>`)
		WriteInt(int64(val), buffer)
		buffer.WriteString(`</li>`)
	}
	buffer.WriteString(`</ul><ul>`)

	for index, val := range []string{"zero", "one", "two"} {
		buffer.WriteString(`<li>`)
		WriteEscString(strconv.Itoa(index)+": "+val, buffer)
		buffer.WriteString(`</li>`)
	}
	buffer.WriteString(`</ul><ul>`)

	for index, val := range map[int]string{1: "one", 2: "two", 3: "three"} {
		buffer.WriteString(`<li>`)
		WriteEscString(strconv.Itoa(index)+": "+val, buffer)
		buffer.WriteString(`</li>`)
	}
	buffer.WriteString(`</ul>`)

	qfs := func(condition bool, iftrue, iffalse []string) []string {
		if condition {
			return iftrue
		} else {
			return iffalse
		}
	}
	var values = []string{}

	buffer.WriteString(`<ul>`)
	for _, val := range qfs(len(values) > 0, values, []string{"There are no values"}) {
		buffer.WriteString(`<li>`)
		WriteEscString(val, buffer)
		buffer.WriteString(`</li>`)
	}
	buffer.WriteString(`</ul>`)
	var values1 = []string{}
	buffer.WriteString(`<ul>`)
	if len(values1) > 0 {
		for _, val := range values1 {
			buffer.WriteString(`<li>`)
			WriteEscString(val, buffer)
			buffer.WriteString(`</li>`)
		}
	} else {
		buffer.WriteString(`<li>There are no values1</li>`)

	}
	buffer.WriteString(`</ul>`)
	var n = 0
	buffer.WriteString(`<ul>`)
	for n < 4 {
		buffer.WriteString(`<li>`)
		WriteInt(int64(n), buffer)
		n++
		buffer.WriteString(`</li>`)
	}
	buffer.WriteString(`</ul>`)

}

// Code generated by "pug.go"; DO NOT EDIT.

package pug

import (
	"strconv"

	pool "github.com/valyala/bytebufferpool"
)

const (
	iteration__0  = `<ul>`
	iteration__1  = `</ul><ul>`
	iteration__3  = `</ul>`
	iteration__20 = `<li>There are no values1</li>`
)

func pug_iteration(buffer *pool.ByteBuffer) {

	buffer.WriteString(iteration__0)
	for _, val := range []int{1, 2, 3, 4, 5} {
		buffer.WriteString(code__6)
		WriteInt(int64(val), buffer)
		buffer.WriteString(code__7)
	}
	buffer.WriteString(iteration__1)

	for index, val := range []string{"zero", "one", "two"} {
		buffer.WriteString(code__6)
		WriteEscString(strconv.Itoa(index)+": "+val, buffer)
		buffer.WriteString(code__7)
	}
	buffer.WriteString(iteration__1)

	for index, val := range map[int]string{1: "one", 2: "two", 3: "three"} {
		buffer.WriteString(code__6)
		WriteEscString(strconv.Itoa(index)+": "+val, buffer)
		buffer.WriteString(code__7)
	}
	buffer.WriteString(iteration__3)

	qfs := func(condition bool, iftrue, iffalse []string) []string {
		if condition {
			return iftrue
		} else {
			return iffalse
		}
	}
	var values = []string{}

	buffer.WriteString(iteration__0)
	for _, val := range qfs(len(values) > 0, values, []string{"There are no values"}) {
		buffer.WriteString(code__6)
		WriteEscString(val, buffer)
		buffer.WriteString(code__7)
	}
	buffer.WriteString(iteration__3)
	var values1 = []string{}
	buffer.WriteString(iteration__0)
	if len(values1) > 0 {
		for _, val := range values1 {
			buffer.WriteString(code__6)
			WriteEscString(val, buffer)
			buffer.WriteString(code__7)
		}
	} else {
		buffer.WriteString(iteration__20)

	}
	buffer.WriteString(iteration__3)
	var n = 0
	buffer.WriteString(iteration__0)
	for n < 4 {
		buffer.WriteString(code__6)
		WriteInt(int64(n), buffer)
		n++
		buffer.WriteString(code__7)
	}
	buffer.WriteString(iteration__3)

}
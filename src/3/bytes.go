package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(buffer([]int{1, 2, 3}))
}

func buffer(values []int) string {
	var b bytes.Buffer

	b.WriteByte(byte('['))
	for index, value := range values {
		// b.WriteByte(byte(value))
		if index > 0 {
			b.WriteString(", ")
		}
		fmt.Fprintf(&b, "%d", value)
	}
	b.WriteByte(byte(']'))

	s := " // hello, 世界"
	for _, v := range []rune(s) {
		b.WriteRune(v)
	}
	return b.String()
}

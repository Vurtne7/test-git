package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	s1 := "abcde abcd"
	s2 := "abcd abcd"
	i := strings.Compare(s1, s2)
	fmt.Println(i)
	fmt.Println(strings.Contains(s1, "abcde"))
	fmt.Println(strings.Count(s1, "d"))

	sliceS1 := strings.Fields(s1)
	fmt.Println(sliceS1)
	fmt.Println(sliceS1[0])
	fmt.Println(sliceS1[1])
	fmt.Println(strings.Join(sliceS1, s2))

	hasPrefix := strings.HasPrefix(s1, "abc")
	fmt.Println(hasPrefix)

	index := strings.Index(s1, "abc")
	fmt.Println(index)
	fmt.Println(strings.LastIndex(s1, "abc"))

	b1 := []byte("abcd abc")
	b2 := []byte("abc")
	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(bytes.Contains(b1, b2))
	fmt.Println(bytes.Count(b1, b2))

	bb1 := bytes.Fields(b1)
	fmt.Println(bb1)
	fmt.Println(bb1[0])
	fmt.Println(bb1[1])

	fmt.Println(bytes.HasPrefix(b1, b2))
	fmt.Println(bytes.Index(b1, b2))

	b3 := bytes.Join(bb1, b2)
	fmt.Println(b3)
	fmt.Println(string(b3))
}

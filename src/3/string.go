package main

import (
	"fmt"
	"unicode/utf8"
)

const useGo = `
go is a tool for managing Go source code.

Usage:
go command [arguments]

\'
\r
\n
...` + "`" + ` xx `

func main() {
	s := "hello, world"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])

	fmt.Printf("%s %s \n", s[0], s[7])
	// fmt.Println(s[len(s)])

	fmt.Println(s[0:5])
	fmt.Println(useGo)

	s1 := "你好世界"
	fmt.Println(len(s1))
	fmt.Println(s1[:6])
	fmt.Println(s1[6:])
	fmt.Println(s1[5:])
	fmt.Println(s1[0], s1[11])
	// cannot assign to s1[0]
	// s1[0] = 229

	s2 := "\x10 \377" // 16进制 8进制
	fmt.Println(s2)

	fmt.Println("\x41")
	fmt.Println(
		"世界",
		"\xe4\xb8\x96\xe7\x95\x8c",
		"\u4e16\u754c",
		"\U00004e16\U0000754c",
	)

	// strings.HasPrefix()
	// strings.HasSuffix()
	// strings.Contains() // 包含子串

	s3 := "Hello, 世界"
	fmt.Println(len(s3))                    // 13
	fmt.Println(utf8.RuneCountInString(s3)) // 9

	for i := 0; i < len(s3); {
		ru, size := utf8.DecodeRuneInString(s3[i:])
		// fmt.Printf("s3[i:]=%s i=%d \t ru=%c \t size=%d \n", s3[i:], i, ru, size)
		fmt.Printf("i=%d \t ru=%c \t size=%d \n", i, ru, size)
		i += size
	}

	n := 0
	for i, r := range s3 {
		fmt.Printf("%d \t %q \t %d \n", i, r, r)
		n++
	}
	fmt.Println(n) // 9
	fmt.Println(utf8.RuneCountInString(s3)) // 9

}

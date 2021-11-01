package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "プログラム"
	fmt.Printf("% x \n", s)
	//fmt.Printf("%x \n", s)

	r := []rune(s)
	fmt.Printf("%x \n", r)
	fmt.Println(string(r))

	fmt.Println(string(65))
	fmt.Println(string(0x4eac))
	fmt.Println(string(12345678))

	fmt.Println(string(int(65)))
	fmt.Println(string("0x4eac"))

	fmt.Println(strings.Compare("hellos", "helloss"))

	fmt.Println(basename("a/b/c.go"))
	fmt.Println(basename2("a/b/c.go"))

	fmt.Println(comma("1234567"))
	fmt.Println(comma2("1234567"))
	fmt.Println(comma2("1234567.789.789"))
	fmt.Println(comma2("12312311234567.78900"))
	fmt.Println(comma2("567.78900"))

	s1 := "abc"
	b := []byte(s1)
	b2 := []rune(s1)
	s2 := string(b)
	s4 := string(b2)
	fmt.Println(s1)
	fmt.Println(b)
	fmt.Println(b2)
	fmt.Println(s2)
	fmt.Println(s4)

	var i int = 12345678
	s3 := string(i)
	fmt.Println(s3)

}

// 可利用 filepath 包
func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s
}

func basename2(s string) string {
	lastIndex := strings.LastIndex(s, "/")
	s = s[lastIndex+1:]

	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	return comma(s[:n-3]) + "," + s[n-3:]
}

func comma2(s string) string {
	var points string
	if strings.Index(s, ".") >= 0 {
		stringSlices := strings.Split(s, ".")
		if len(stringSlices) >= 3 {
			// ERROR
			return s
		}
		s = stringSlices[0]
		points = "." + stringSlices[1]
	}

	n := len(s)
	if n <= 3 {
		return s + points
	}

	return comma2(s[:n-3]) + "," + s[n-3:] + points
}

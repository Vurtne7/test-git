package main

import (
	"fmt"
)

func main() {
	fmt.Println(compare("eaddadbcffkjhlahflahpfqjfqlifqfc", "cba"))
	//fmt.Println(compare("abcd", "dcba"))
	//fmt.Println(compare("abcd", "dcbe"))
	//fmt.Println(compare("abcda", "dacba"))
}

func compare(s1, s2 string) bool {
	if len(s1) != len(s2) {
		// return false
	}

	b1, b2 := []byte(s1), []byte(s2)
	b1 = sortSlice(b1)
	b2 = sortSlice(b2)
	fmt.Println(b1)
	fmt.Println(b2)

	/*
		sort.Slice(b1, func(i, j int) bool {
			return b1[i] < b1[j]
		})
		sort.Slice(b2, func(i, j int) bool {
			return b2[i] < b2[j]
		})
	*/

	r1, r2 := []rune(s1), []rune(s2)
	r1 = sortSlice2(r1)
	r2 = sortSlice2(r2)
	fmt.Println(r1)
	fmt.Println(r2)
	/*
		sort.Slice(r1, func(i, j int) bool {
			return r1[i] > r1[j]
		})
		sort.Slice(r2, func(i, j int) bool {
			return r2[i] > r2[j]
		})
	*/

	return string(b1) == string(b2) && string(r1) == string(r2)
}

func sortSlice(b []byte) []byte {
	for i := 0; i < len(b); i++ {
		for j := i + 1; j < len(b); j++ {
			if b[i] > b[j] {
				b[i], b[j] = b[j], b[i]
			}
		}
	}
	return b
}

func sortSlice2(b []rune) []rune {
	if len(b) <= 1 {
		return b
	}

	var leftB, rightB []rune
	start := b[0]
	b = b[1:]
	for _, v := range b {
		if v <= start {
			leftB = append(leftB, v)
		} else {
			rightB = append(rightB, v)
		}
	}

	if len(leftB) > 0 {
		leftB = sortSlice2(leftB)
	}
	if len(rightB) > 0 {
		rightB = sortSlice2(rightB)
	}

	leftB = append(leftB, start)
	return append(leftB, rightB...)
}

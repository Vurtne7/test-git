package main

import (
	"fmt"
)

// ERROR
//var a int = b
//var b int = a

var pc1 [256]byte = func() (pc1 [256]byte) {
	for i := range pc1 {
		pc1[i] = pc1[i/2] + byte(i&1)
	}
	return
}()

func PopCount1(x uint64) int {
	result, byteV := 0, 8
	for i := 0; i < byteV; i++ {
		result += int(pc1[byte(x>>(i*byteV))])
	}
	return result
}

func main() {
	a := 1
	fmt.Println(a)
	a = func() (b int) {
		// var b int = 0
		fmt.Println(b)
		b = 3
		return // b is already = 3
	}()
	fmt.Println(a)

	fmt.Println(PopCount1(7))

	x := "hello!"
	for i := 0; i < len(x); i++ {
		x := x[i]
		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("%c", x) // "HELLO" (one letter per iteration)
		}
	}
}

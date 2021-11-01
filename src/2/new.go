package main

import (
	"fmt"
)

func main() {
	p := new(struct{})
	fmt.Println(*p)

	p2 := new([0]int)
	fmt.Println(*p2)
}

func delta(new, old int) int {
	// ERROR
	// p := new(int)
	// fmt.Println(*p)
	return new - old
}

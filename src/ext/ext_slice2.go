package main

import "fmt"

func main() {
	a := []int{2, 4, 3, 4}
	b := []int{1, 2}
	ints := copy(a, b)

	fmt.Println(ints)
	fmt.Println(a)
	fmt.Println(b)
}

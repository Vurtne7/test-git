package main

import (
	"fmt"
)

func main() {
	a := [...]int{1, 2, 3}
	s := a[0:1]
	fmt.Println(a)
	fmt.Println(s)

	s[0] = 5
	fmt.Println(a)
	fmt.Println(s)

	a[0] = 55
	fmt.Println(a)
	fmt.Println(s)
}

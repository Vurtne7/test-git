package main

import "fmt"

func main() {
	var p *int
	fmt.Println(p)
	a := 1
	p = &a
	fmt.Println(p)
	fmt.Println(*p)


	var x int = 2
	*p = x
	fmt.Println(a)
	fmt.Println(*p)


	var p2 *int
	x2 := new(int)
	p2 = x2
	fmt.Println(*p2)
}

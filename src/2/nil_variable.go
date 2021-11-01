package main

import "fmt"

func main() {
	var p *int
	p = nil
	fmt.Println(p)

	slices := []string{}
	slices = nil
	fmt.Println(slices)

	maps := map[int]string{}
	maps = nil
	fmt.Println(maps)

	var chans chan string
	chans = nil
	fmt.Println(chans)


	type a int

	var n a = 1
	fmt.Println(n)

	var ns struct{}
	fmt.Println(ns)
}

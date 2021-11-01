package main

import "fmt"

func main() {

	arr := [...]int{1, 2, 3}
	fmt.Println(arr)
	changeArray(arr)
	fmt.Println(arr)
}

func changeArray(a [3]int) {
	a[0] = 10
	fmt.Println(a)
}

package main

import "fmt"

func main() {
	var j *int
	fmt.Println(j)

	//b := 10
	//j = &b
	//fmt.Println(*j)

	j = new(int)
	fmt.Println(*j)
	*j = 10
	fmt.Println(*j)

	var slices []int
	if slices == nil {
		fmt.Println("slices is nil 1")
	}
	fmt.Println(slices)
	slices = make([]int, 5, 5) // 长度=5 容量=5
	if slices == nil {
		fmt.Println("slices is nil 2")
	}
	fmt.Println(slices)
	fmt.Println(len(slices))
	fmt.Println(cap(slices))
	slices = make([]int, 10, 15)
	fmt.Println(slices[len(slices) - 1])
	fmt.Println(slices)
	changeSlice(slices)
	fmt.Println(slices)

	// ERROR
	// var slices2 []int
	// slices2[0] = 1
	// fmt.Println(slices2)
	// slices2 = make([]int, 10)
	// changeSlice(slices2)

	// ERROR
	// var map2 map[string]string
	// map2["123"] = "123"
	// fmt.Println(map2)


	aa := [3]int {1, 2, 3}
	fmt.Println(aa[len(aa) - 1])
	fmt.Println(cap(aa))

	//b := make([...]int, 10)
	//fmt.Println(b)
	//bb := [...]int {1, 2, 3}

	// ERROR
	// var slices2 []int
	// slices2 := []int{}
	// changeSlice(slices2)
	// fmt.Println(slices2)

	var maps map[string]int
	if maps == nil {
		fmt.Println("maps is nil 1")
	}
	fmt.Println(maps)
	maps = make(map[string]int, 10)
	if maps == nil {
		fmt.Println("maps is nil 2")
	}
	fmt.Println(maps)
	changeMap(maps)
	fmt.Println(maps)

	// new 和 make都是Go语言的两个内建函数，用于分配内存
	// new 一般用来返回指针类型（一般不用），make返回引用类型（map, slice,chan 这三个引用)
	// var 声明的 基本类型和struct这种已经分配了内存，并且赋零值了。
}

func changeSlice(s []int) {
	//s = append(s, 10)
	s[0] = 10
	s[9] = 10
	s = append(s, 10)
	fmt.Println(s[10])
	fmt.Println(s)
}

func changeMap(m map[string]int) {
	m["1234"] = 1234
	fmt.Println(m)
}

package main

import (
	"fmt"
	"math"
)

func main() {
	var a int8
	var b byte
	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)

	var c rune
	var d int32
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)

	var ff = 10
	var fff = +ff
	var ffff = -ff

	ff &^= 3
	fmt.Println(ff)
	fmt.Println(fff)
	fmt.Println(ffff)

	ascii := 'a'
	unicode := '国'
	unicode2 := "国"
	newline := '\n'

	fmt.Printf("%d %[1]c %[1]q \n", ascii)
	fmt.Printf("%d %[1]c %[1]q \n", unicode)
	fmt.Printf("%d %[1]c %[1]q \n", unicode2)
	fmt.Printf("%d %[1]q \n", newline)

	const e = .77
	const e1 = 1.

	for i := 0; i < 8; i++ {
		fmt.Printf("x = %d  e^x = %8.3f \n", i, math.Exp(float64(i)))
	}

	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z) // "0 -0 +Inf -Inf NaN"

	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan > nan) // "false false false"

	if math.IsNaN(nan) {
		fmt.Println("is nan")
	}

	fmt.Println(
		math.IsInf(1/z, 1),   // is +inf
		math.IsInf(-1/z, -1), // is -inf
		math.IsInf(1/z, 0),   // is inf
		math.IsInf(-1/z, 0),  // is inf
	) // all true
}

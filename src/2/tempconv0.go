package main

import (
	"fmt"
)

type Celsius float64    // 摄氏温度
type Fahrenheit float64 // 华氏温度

const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 结冰点温度
	BoilingC      Celsius = 100     // 沸水温度
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

type Points *int

func Po(i int) Points {
	// (*int)(&i)
	return Points(&i)
}

func main() {
	fmt.Printf("%g \n", BoilingC-FreezingC) // 100
	BoilingF := CToF(BoilingC)
	fmt.Printf("%g \n", BoilingF-CToF(FreezingC)) // 180
	// fmt.Printf("%g \n", BoilingF-FreezingC) // invalid operation: BoilingF - FreezingC (mismatched types Fahrenheit and Celsius)

	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0) // true
	fmt.Println(f >= 0) // true
	// fmt.Println(c == f) // invalid operation: c == f (mismatched types Celsius and Fahrenheit)
	fmt.Println(c == Celsius(f)) // true

	c = FToC(212.0)
	fmt.Println(c.String())			// "100°C"
	fmt.Printf("%v \n", c)	// "100°C"; no need to call String explicitly
	fmt.Printf("%s \n", c)	// "100°C"
	fmt.Println(c)					// "100°C"
	fmt.Printf("%g \n", c)	// "100"; does not call String
	fmt.Println(float64(c)) 		// "100"; does not call String
}

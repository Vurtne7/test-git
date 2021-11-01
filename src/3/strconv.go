package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 9
	fmt.Println(strconv.FormatInt(int64(x), 2))

	y := fmt.Sprintf("%d", x)
	fmt.Println(y)
	fmt.Println(strconv.Itoa(x))

	s := fmt.Sprintf("x=%b", x)
	fmt.Println(s)

	i, _ := strconv.Atoi("123")
	fmt.Println(i)

	j, _ := strconv.ParseInt("123", 10, 64)
	fmt.Println(j)
}

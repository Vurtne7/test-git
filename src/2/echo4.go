package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "n desc")
var s = flag.String("s", " ", "s desc")

func main() {
	flag.Parse()

	fmt.Println(strings.Join(flag.Args(), *s))
	if n == nil {
		fmt.Println("n is nil")
	}
	if *n == false {
		fmt.Println("n is false")
	}
	fmt.Println(n)
	fmt.Println(*n)
}

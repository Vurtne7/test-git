package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()

	fmt.Println(os.Args[0])

	s, sep := "", ""
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	secs := time.Since(start).Milliseconds()
	fmt.Printf("%ds \n", secs)
}

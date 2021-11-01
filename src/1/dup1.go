package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	inputs := bufio.NewScanner(os.Stdin)

	flag := 0
	for inputs.Scan() {
		counts[inputs.Text()]++
		flag++
		if flag > 10 {
			break
		}
	}

	for line, i := range counts {
		if i > 1 {
			fmt.Printf("string: %s, number: %d \n", line, i)
		}
	}
}

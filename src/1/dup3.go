package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	files := os.Args[1:]
	counts := make(map[string]int)

	for _, file := range files {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Printf("ioutil read file failed, filename: %s, error: %v \n", file, err)
			continue
		}

		stringSlices := strings.Split(string(data), "\n")
		for _, line := range stringSlices {
			counts[line]++
		}
	}

	for line, number := range counts {
		if number > 1 {
			fmt.Printf("line: %s, number: %d \n", line, number)
		}
	}
}

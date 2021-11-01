package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// var counts map[string]int
	var counts map[string]map[string]int
	counts = make(map[string]map[string]int)
	fmt.Println(counts)

	args := os.Args[1:]
	if len(args) == 0 {
		contentLines(os.Stdin, counts)
	} else {
		for _, arg := range args {
			file, err := os.Open(arg)
			if err != nil {
				fmt.Printf("file: %s is not exists, err: %v \n", arg, err)
				continue
			}
			contentLines(file, counts)
			file.Close()
		}
	}

	for filename, lineData := range counts {
		for line, i := range lineData {
			if i > 1 {
				fmt.Printf("filename %s, %s, %d \n", filename, line, i)
			}
		}
	}
}

func contentLines(file *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		if counts[file.Name()] == nil {
			counts[file.Name()] = make(map[string]int)
		}
		counts[file.Name()][input.Text()]++
	}
}

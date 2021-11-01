package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {

	// 获取命令行参数1
	flag.Parse()
	if len(flag.Args()) > 1 {
		for _, arg := range flag.Args() {
			fmt.Println(arg)
		}
		return
	}

	// 获取命令行参数2
	if len(os.Args) > 2 {
		for _, arg := range os.Args[1:] {
			fmt.Println(arg)
		}
		return
	}

	// 获取输入的参数
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

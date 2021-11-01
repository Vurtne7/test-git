package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("empty url \n")
		os.Exit(1)
	}

	url := os.Args[1]
	if url == "" {
		fmt.Printf("empty url \n")
		os.Exit(1)
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("http get error: %v \n", err)
		os.Exit(1)
	}

	data, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Printf("get response body error: %v \n", err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", data)
	os.Exit(1)
}

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("empty urls \n")
		os.Exit(1)
	}

	for _, url := range os.Args[1:] {
		if url == "" {
			fmt.Printf("empty url \n")
			os.Exit(1)
		}
		if strings.HasPrefix(url, "http") == false {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("http get error: %v \n", err)
			os.Exit(1)
		}
		if resp.StatusCode != http.StatusOK {
			fmt.Printf("http status no ok: %d \n", resp.StatusCode)
			os.Exit(1)
		}

		// bytes, err := ioutil.ReadAll(resp.Body)
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Printf("read all error: %v \n", err)
			os.Exit(1)
		}

		// fmt.Printf("%v \n", w)
	}
}

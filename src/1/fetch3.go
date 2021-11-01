package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	if len(os.Args) < 2 {
		fmt.Printf("empty urls \n")
		os.Exit(1)
	}

	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		printLn(ch)
	}
	secs := time.Since(start).Seconds()
	fmt.Printf("total: %.2fs \n", secs)
}

func printLn(ch <-chan string) {
	fmt.Println(<-ch)
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	if url == "" {
		ch <- fmt.Sprintf("url is empty string")
		return
	}
	if strings.HasPrefix(url, "http") == false {
		url = "http://" + url
	}

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("%s, http get error: %v", url, err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		ch <- fmt.Sprintf("%s, http status code not ok: %d", url, resp.StatusCode)
		return
	}

	written, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("%s, io copy error: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, written, url)
	return
}

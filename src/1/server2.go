package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", HandleRoot)
	http.HandleFunc("/counts", HandleCount)
	log.Fatal(http.ListenAndServe("localhost:9777", nil))
}

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s \n", r.Method, r.URL, r.Proto)

	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q]=%q \n", k, v)
	}

	fmt.Fprintf(w, "Host=%q \n", r.Host)
	fmt.Fprintf(w, "RemoteAddr=%q \n", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		log.Println(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q]=%q \n", k, v)
	}

	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "Url.Path=%q", r.URL.Path)
}

func HandleCount(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "count: %d \n", count)
	mu.Unlock()
}

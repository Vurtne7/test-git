package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleFun)
	log.Fatal(http.ListenAndServe("localhost:9777", nil))
}

func handleFun(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Url.Path=%q\n", r.URL.Path)
}

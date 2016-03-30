package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Host == "peterrosser.com" {
		fmt.Fprintf(w, "Hello from pr")
	} else if r.Host == "thefirsttrust.org" {
		fmt.Fprintf(w, "Hello from tft")
	} else {
		fmt.Fprintf(w, "Hello from somewhere else")
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
	http.ListenAndServe(":8080", nil)
}

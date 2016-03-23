package main

import (
	"fmt"
	"net"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	host, _, _ := net.SplitHostPort(r.Host)
	fmt.Fprintf(w, "Hello from %s", host)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
	http.ListenAndServe(":8080", nil)
}

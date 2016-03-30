package main

import (
	"fmt"
	"net"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	host, port, _ := net.SplitHostPort(r.Host)
	fmt.Fprintf(w, "Hello from %s on port %s, %s", host, port, r.Host)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"fmt"
	"net"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	host, port, _ := net.SplitHostPort(r.Host)
	fmt.Fprintf(w, "Hello from %s on port %s", host, port)
}

func main() {
	http.ListenAndServe(":80", nil)
}

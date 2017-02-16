package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func addRosserSoftwareHost(r *mux.Router) {
	s := r.Host("rosser.software").Subrouter()
	s.HandleFunc("/helloworld", HelloWorldHandler)
}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Well, hello to you too! thanks for visiting %s", r.URL.Path[1:])
}

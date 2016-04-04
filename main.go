package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Host == "peterrosser.com" {
		ServePeterRosserWebsite()
		//fmt.Fprintf(w, "Hello from pr")
	} else if r.Host == "thefirsttrust.org" {
		ServeTheFirstTrustWebsite()
		//fmt.Fprintf(w, "Hello from tft")
	} else {
		fmt.Fprintf(w, "Hello from somewhere else")
	}
}

func main() {
	http.HandleFunc("/", handler)
}

func ServePeterRosserWebsite() {
	log.Fatal(http.ListenAndServe(":80", http.FileServer(http.Dir("/peterrosser"))))
}

func ServeTheFirstTrustWebsite() {
	log.Fatal(http.ListenAndServe(":80", http.FileServer(http.Dir("/thefirsttrust"))))
}

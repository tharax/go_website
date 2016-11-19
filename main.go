package main

import "net/http"

func main() {
	http.Handle("peterrosser.com/", http.FileServer(http.Dir("./peterrosser")))
	http.Handle("thefirsttrust.org/", http.FileServer(http.Dir("./thefirsttrust")))
	http.Handle("localhost/", http.FileServer(http.Dir("./localhost")))
	http.ListenAndServe(":80", nil)
}
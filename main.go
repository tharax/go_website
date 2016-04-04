package main

import "net/http"

func main() {
	http.ListenAndServe(":80", http.FileServer(http.Dir("./peterrosser")))
	http.ListenAndServe(":8080", http.FileServer(http.Dir("./peterrosser")))
}

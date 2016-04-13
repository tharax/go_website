package main

import "net/http"

func main() {
	mux := http.NewServeMux()

	rhpr := http.RedirectHandler("peterrosser.com", 8080)
	mux.Handle("/peterrosser.com", rhpr)
	rhtft := http.RedirectHandler("thefirsttrust.org", 8080)
	mux.Handle("/thefirsttrust.com", rhtft)
	rhlh := http.RedirectHandler("localhost", 8080)
	mux.Handle("/localhost", rhlh)

	http.ListenAndServe(":8080", mux)
}

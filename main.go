package main

import "net/http"

func main() {
	mux := http.NewServeMux()

	rhpr := http.RedirectHandler("peterrosser.com", 8080)
	mux.Handle("/peterrosser", rhpr)
	rhtft := http.RedirectHandler("thefirsttrust.org", 8080)
	mux.Handle("/thefirsttrust", rhtft)
	rhlh := http.RedirectHandler("localhost", 8080)
	mux.Handle("/localhost", rhlh)

	http.ListenAndServe(":8080", mux)
}

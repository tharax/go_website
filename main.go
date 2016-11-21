package main

import "net/http"

func main() {

	//This is only here for testing.
	http.Handle("localhost/", http.FileServer(http.Dir("./localhost")))

	//This adds additional handlers to the default mux.
	http.Handle("peterrosser.com/", http.FileServer(http.Dir("./peterrosser")))
	http.Handle("thefirsttrust.org/", http.FileServer(http.Dir("./thefirsttrust")))
	http.Handle("rosser.software/", http.FileServer(http.Dir("./rossersoftware")))
	http.Handle("rossersoftware.com/", http.FileServer(http.Dir("./rossersoftware")))

	//By using "nil" we use the default mux.
	http.ListenAndServeTLS(":443", "/etc/letsencrypt/live/peterrosser.com/cert.pem", "/etc/letsencrypt/live/peterrosser.com/privkey.pem", nil)
}
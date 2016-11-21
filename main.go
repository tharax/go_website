package main

import "net/http"

func main() {

	// redirect every http request to https
	go http.ListenAndServe(":80", http.HandlerFunc(redirect))

	//This is only here for testing.
	http.Handle("localhost/", http.FileServer(http.Dir("./localhost")))

	//This adds additional handlers to the default mux.
	http.Handle("https://peterrosser.com/", http.FileServer(http.Dir("./peterrosser")))
	http.Handle("thefirsttrust.org/", http.FileServer(http.Dir("./thefirsttrust")))
	http.Handle("rosser.software/", http.FileServer(http.Dir("./rossersoftware")))
	http.Handle("rossersoftware.com/", http.FileServer(http.Dir("./rossersoftware")))

	//By using "nil" we use the default mux.
	http.ListenAndServeTLS(":443", "/etc/letsencrypt/live/peterrosser.com/cert.pem", "/etc/letsencrypt/live/peterrosser.com/privkey.pem", nil)
}

func redirect(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "https://"+req.Host+req.URL.String(), http.StatusMovedPermanently)
}

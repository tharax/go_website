package main

import "net/http"

func main() {

	// redirect every http request to https
	go http.ListenAndServe(":80", http.HandlerFunc(redirect))

	startSimpleServer("peterrosser.com", "./peterrosser")
	//go startSimpleServer("thefirsttrust.org", "./thefirsttrust")
	//go startSimpleServer("rosser.software", "./rossersoftware")
	//go startSimpleServer("rossersoftware.com", "./rossersoftware")
}

func redirect(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "https://"+req.Host+req.URL.String(), http.StatusMovedPermanently)
}

func startSimpleServer(serverName, serverFolder string) {
	//This adds additional handlers to the default mux.
	http.Handle("/", http.FileServer(http.Dir(serverFolder)))
	//By using "nil" we use the default mux.
	http.ListenAndServeTLS(":443", "/etc/letsencrypt/live/"+serverName+"/cert.pem", "/etc/letsencrypt/live/"+serverName+"/privkey.pem", nil)
}

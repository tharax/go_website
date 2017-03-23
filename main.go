package main

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tharax/websites"
)

func main() {

	// redirect every http request to https.
	go http.ListenAndServe(":80", http.HandlerFunc(redirectToTLS))

	// create a list of all SSL certs.
	config := &tls.Config{}
	config.Certificates = append(config.Certificates, getCert("peterrosser.com"))
	config.Certificates = append(config.Certificates, getCert("thefirsttrust.org"))
	config.Certificates = append(config.Certificates, getCert("rosser.software"))
	config.Certificates = append(config.Certificates, getCert("rossersoftware.com"))
	config.BuildNameToCertificate()

	// create different handlers for different hosts.
	r := mux.NewRouter()
	r.Host("peterrosser.com").Handler(websites.PersonalHandler())
	r.Host("thefirsttrust.org").Handler(websites.TrustHandler())
	r.Host("rossersoftware.com").Handler(websites.BusinessHandler())
	r.Host("rosser.software").Handler(websites.BusinessHandler())

	// create the server.
	server := http.Server{
		Addr:      ":443",
		Handler:   r,
		TLSConfig: config,
	}

	server.ListenAndServeTLS("", "")

}

func getCert(website string) (cert tls.Certificate) {
	cert, err := tls.LoadX509KeyPair("/etc/letsencrypt/live/"+website+"/cert.pem", "/etc/letsencrypt/live/"+website+"/privkey.pem")
	if err != nil {
		log.Fatal(err)
	}
	return cert
}

func redirectToTLS(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "https://"+req.Host+req.URL.String(), http.StatusMovedPermanently)
}

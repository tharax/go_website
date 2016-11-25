package main

import "net/http"
import "crypto/tls"
import "log"

func main() {

	// redirect every http request to https
	go http.ListenAndServe(":80", http.HandlerFunc(redirect))
	go http.ListenAndServe("www", http.HandlerFunc(redirect))

	cfg := &tls.Config{}

	cert, err := tls.LoadX509KeyPair("/etc/letsencrypt/live/"+"peterrosser.com"+"/cert.pem", "/etc/letsencrypt/live/"+"peterrosser.com"+"/privkey.pem")
	if err != nil {
		log.Fatal(err)
	}
	cfg.Certificates = append(cfg.Certificates, cert)
	cert, err = tls.LoadX509KeyPair("/etc/letsencrypt/live/"+"thefirsttrust.org"+"/cert.pem", "/etc/letsencrypt/live/"+"thefirsttrust.org"+"/privkey.pem")
	if err != nil {
		log.Fatal(err)
	}
	cfg.Certificates = append(cfg.Certificates, cert)
	cert, err = tls.LoadX509KeyPair("/etc/letsencrypt/live/"+"rosser.software"+"/cert.pem", "/etc/letsencrypt/live/"+"rosser.software"+"/privkey.pem")
	if err != nil {
		log.Fatal(err)
	}
	cfg.Certificates = append(cfg.Certificates, cert)
	cert, err = tls.LoadX509KeyPair("/etc/letsencrypt/live/"+"rossersoftware.com"+"/cert.pem", "/etc/letsencrypt/live/"+"rossersoftware.com"+"/privkey.pem")
	if err != nil {
		log.Fatal(err)
	}
	cfg.Certificates = append(cfg.Certificates, cert)

	cfg.BuildNameToCertificate()

	server := http.Server{
		Addr:      ":443",
		Handler:   http.FileServer(http.Dir("./peterrosser")),
		TLSConfig: cfg,
	}

	server.ListenAndServeTLS("", "")

}

func redirect(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "https://"+req.Host+req.URL.String(), http.StatusMovedPermanently)
}

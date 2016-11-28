package main

import "net/http"
import "crypto/tls"
import "log"

func main() {

	// redirect every http request to https.
	go http.ListenAndServe(":80", http.HandlerFunc(redirect))

	cfg := &tls.Config{}
	cfg.Certificates = append(cfg.Certificates, getCert("peterrosser.com"))
	cfg.Certificates = append(cfg.Certificates, getCert("thefirsttrust.org"))
	cfg.Certificates = append(cfg.Certificates, getCert("rosser.software"))
	cfg.Certificates = append(cfg.Certificates, getCert("rossersoftware.com"))
	cfg.BuildNameToCertificate()

	server := http.Server{
		Addr:      ":443",
		Handler:   http.FileServer(http.Dir("./peterrosser")),
		TLSConfig: cfg,
	}

	server.ListenAndServeTLS("", "")

}

func getCert(website string) (cert Certificate) {
	cert, err = tls.LoadX509KeyPair("/etc/letsencrypt/live/"+website+"/cert.pem", "/etc/letsencrypt/live/"+website+"/privkey.pem")
	if err != nil {
		log.Fatal(err)
	}
	return cert
}

func redirect(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "https://"+req.Host+req.URL.String(), http.StatusMovedPermanently)
}

package main

import "net/http"
import "crypto/tls"
import "log"
import "fmt"

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
		Handler:   http.HandlerFunc(getHandler),
		TLSConfig: cfg,
	}

	server.ListenAndServeTLS("", "")

}

func getHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from %s ", r.URL.Host)
	// if r.Host == "peterrosser.com" {
	// 	http.FileServer(http.Dir("./peterrosser"))
	// } else if r.Host == "thefirsttrust.org" {
	// 	http.FileServer(http.Dir("./thefirsttrust"))
	// } else if r.Host == "rosser.software" {
	// 	http.FileServer(http.Dir("./rossersoftware"))
	// } else if r.Host == "rossersoftware.com" {
	// 	http.FileServer(http.Dir("./rossersoftware"))
	// } else {
	// 	http.NotFoundHandler()
	// }
}

func getCert(website string) (cert tls.Certificate) {
	cert, err := tls.LoadX509KeyPair("/etc/letsencrypt/live/"+website+"/cert.pem", "/etc/letsencrypt/live/"+website+"/privkey.pem")
	if err != nil {
		log.Fatal(err)
	}
	return cert
}

func redirect(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "https://"+req.Host+req.URL.String(), http.StatusMovedPermanently)
}

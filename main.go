package main

import "net/http"
import "crypto/tls"
import "log"
import "github.com/gorilla/mux"

func main() {

	// redirect every http request to https.
	go http.ListenAndServe(":80", http.HandlerFunc(redirect))

	cfg := &tls.Config{}
	cfg.Certificates = append(cfg.Certificates, getCert("peterrosser.com"))
	cfg.Certificates = append(cfg.Certificates, getCert("thefirsttrust.org"))
	cfg.Certificates = append(cfg.Certificates, getCert("rosser.software"))
	cfg.Certificates = append(cfg.Certificates, getCert("rossersoftware.com"))
	cfg.BuildNameToCertificate()

	// mux := http.NewServeMux()
	// mux.HandleFunc("peterrosser.com", peterRosserHandler)
	// mux.HandleFunc("thefirsttrust.org", theFirstTrustHandler)
	// mux.HandleFunc("rossersoftware.com", rosserSoftwareHandler)
	// mux.HandleFunc("rosser.software", rosserSoftwareHandler)

	r := mux.NewRouter()
	r.Host("peterrosser.com").Handler(http.FileServer(http.Dir("./peterrosser")))
	r.Host("thefirsttrust.org").Handler(http.FileServer(http.Dir("./thefirsttrust")))

	server := http.Server{
		Addr:      ":443",
		Handler:   r,
		TLSConfig: cfg,
	}

	server.ListenAndServeTLS("", "")

}

// func peterRosserHandler(w http.ResponseWriter, r *http.Request) {
// 	w = http.FileServer(http.Dir("./peterrosser"))
// }

// func theFirstTrustHandler(w http.ResponseWriter, r *http.Request) {
// 	w = http.FileServer(http.Dir("./thefirsttrust"))
// }

// func rosserSoftwareHandler(w http.ResponseWriter, r *http.Request) {
// 	w = http.FileServer(http.Dir("./rossersoftware"))
// }

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

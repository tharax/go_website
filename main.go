package main

import "net/http"
import "crypto/tls"
import "log"

func main() {

	// redirect every http request to https
	go http.ListenAndServe(":80", http.HandlerFunc(redirect))

	ps := peterServer{"peterrosser.com", "./peterrosser"}
	go startServer(ps)

	tft := peterServer{"thefirsttrust.org", "./thefirsttrust"}
	go startServer(tft)

	// startSimpleServer("thefirsttrust.org", "./thefirsttrust")
	//go startSimpleServer("rosser.software", "./rossersoftware")
	//go startSimpleServer("rossersoftware.com", "./rossersoftware")
}

func redirect(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "https://"+req.Host+req.URL.String(), http.StatusMovedPermanently)
}

// func startSimpleServer(serverName, serverFolder string) {
// 	server := http.Server()
// 	server.Addr = ":443"
// 	server.Handler = http.FileServer(http.Dir(serverFolder))
// 	server.TLSConfig = tls.Config()

// 	cert, err := tls.LoadX509KeyPair("/etc/letsencrypt/live/"+serverName+"/cert.pem", "/etc/letsencrypt/live/"+serverName+"/privkey.pem")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	cfg.Certificates = append(cfg.Certificates, cert)
// 	// keep adding remaining certs to cfg.Certificates

// 	cfg.BuildNameToCertificate()

// 	//This adds additional handlers to the default mux.
// 	http.Handle("/", http.FileServer(http.Dir(serverFolder)))
// 	//By using "nil" we use the default mux.
// 	http.ListenAndServeTLS(":443", "/etc/letsencrypt/live/"+serverName+"/cert.pem", "/etc/letsencrypt/live/"+serverName+"/privkey.pem", nil)
// }

type peterServer struct {
	serverName   string
	serverFolder string
}

func startServer(ps peterServer) {
	cfg := &tls.Config{}

	cert, err := tls.LoadX509KeyPair("/etc/letsencrypt/live/"+ps.serverName+"/cert.pem", "/etc/letsencrypt/live/"+ps.serverName+"/privkey.pem")
	if err != nil {
		log.Fatal(err)
	}

	cfg.Certificates = append(cfg.Certificates, cert)
	// keep adding remaining certs to cfg.Certificates

	cfg.BuildNameToCertificate()

	server := http.Server{
		Addr:      ":443",
		Handler:   http.FileServer(http.Dir(ps.serverFolder)),
		TLSConfig: cfg,
	}

	server.ListenAndServeTLS("", "")
}

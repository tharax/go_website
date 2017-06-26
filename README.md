# webserver
A simple program that can host multiple sites from a single server, which translates to cheaper hosting costs.

[![Go Report Card](https://goreportcard.com/badge/github.com/tharax/webserver)](https://goreportcard.com/report/github.com/tharax/webserver)

Features:
* Redirects all http requests to https.
* Can serve multiple TLS certs.
* Extensible - just add a new handler. I add mine as separate libraries.

Wishlist:

* Logging.
* Recover from crashes.
* Metrics.

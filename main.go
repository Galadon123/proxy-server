package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func main() {
	targets := map[string]*url.URL{
		"ns1": parseURL("http://localhost:5001"),
		"ns2": parseURL("http://localhost:5002"),
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		segments := strings.SplitN(r.URL.Path, "/", 3)
		if len(segments) < 2 {
			http.NotFound(w, r)
			return
		}

		targetURL, ok := targets[segments[1]]
		if !ok {
			http.NotFound(w, r)
			return
		}

		proxy := httputil.NewSingleHostReverseProxy(targetURL)
		proxy.ServeHTTP(w, r)
	})

	// Log that the server is running
	log.Println("Server is running on port 3000...")

	// Start the server
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}

func parseURL(rawURL string) *url.URL {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}
	return parsedURL
}

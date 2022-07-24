package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/toivjon/max-pondus/backend/internal/server/admin"
	"github.com/toivjon/max-pondus/backend/internal/server/personal"
)

func main() {
	port := flag.Int("port", 8080, "The port to listen for the incoming HTTP connections.")
	timeout := flag.Duration("timeout", time.Second, "The timeout for processing the request.")
	flag.Parse()

	mux := http.NewServeMux()
	mux.Handle("/api/v1/personal", personal.NewHandler())
	mux.Handle("/api/v1/admin", admin.NewHandler())
	mux.Handle("/", http.NotFoundHandler())

	server := &http.Server{
		// Just use the default hostname and only specify the port we want to listen.
		Addr: fmt.Sprintf(":%d", *port),
		// Make sure that our response writer has time to write the response when processing timeouts.
		WriteTimeout: *timeout + time.Second,
		// Reserve one second to read the request payload.
		ReadTimeout: time.Second,
		// Reserve one second to request to establish a connection and us to read the headers.
		ReadHeaderTimeout: time.Second,
		// Use timeout handler to automatically abort the request handling if the processing timeouts.
		Handler: http.TimeoutHandler(mux, *timeout, ""),
	}
	log.Printf("Starting a server port: %d timeout: %v", *port, *timeout)
	log.Fatal(server.ListenAndServe())
}

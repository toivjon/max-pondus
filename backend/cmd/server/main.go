package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/toivjon/max-pondus/backend/internal/server"
)

func main() {
	port := flag.Int("port", 8080, "The port to listen for the incoming HTTP connections.")
	timeout := flag.Duration("timeout", time.Second, "The timeout for processing the request.")
	flag.Parse()

	mux := http.NewServeMux()
	mux.Handle("/", &server.Handler{})

	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", *port),
		WriteTimeout:      *timeout + time.Second,
		ReadHeaderTimeout: time.Second,
		Handler:           http.TimeoutHandler(mux, *timeout, ""),
	}
	log.Printf("Starting a server port: %d timeout: %v", *port, *timeout)
	log.Fatal(server.ListenAndServe())
}

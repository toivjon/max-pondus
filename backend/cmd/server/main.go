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
	timeout := flag.Int("timeout", 1, "The timeout for the request handling in seconds.")
	flag.Parse()

	mux := http.NewServeMux()
	mux.Handle("/", &server.Handler{})

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", *port),
		WriteTimeout: time.Duration(*timeout+1) * time.Second,
		Handler:      http.TimeoutHandler(mux, time.Duration(*timeout)*time.Second, ""),
	}
	log.Printf("Starting a server at port %d with %d sec timeout", *port, *timeout)
	log.Fatal(server.ListenAndServe())
}

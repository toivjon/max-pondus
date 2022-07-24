package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/toivjon/max-pondus/backend/internal/server"
)

func main() {
	port := flag.Int("port", 8080, "The port to listen for the incoming HTTP connections.")
	flag.Parse()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", *port),
		Handler: &server.Handler{},
	}
	log.Printf("Starting a server at port %d", *port)
	log.Fatal(server.ListenAndServe())
}

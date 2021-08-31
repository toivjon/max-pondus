package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/toivjon/max-pondus/backend/internal/server/health"
	"github.com/toivjon/max-pondus/backend/internal/server/muscle"
	"github.com/toivjon/max-pondus/backend/internal/server/user"
)

func main() {
	port := flag.Int("port", 8080, "The port to listen for the incoming HTTP connections.")
	flag.Parse()

	http.Handle("/health", &health.Handler{})
	http.Handle("/muscles", &muscle.Handler{})
	http.Handle("/users", &user.Handler{})

	log.Printf("Starting a server at port %d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}

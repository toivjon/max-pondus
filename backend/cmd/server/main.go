package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/toivjon/max-pondus/backend/internal/server"
	"github.com/toivjon/max-pondus/backend/internal/server/admin/user"
)

func main() {
	port := flag.Int("port", 8080, "The port to listen for the incoming HTTP connections.")
	flag.Parse()

	app := server.NewApplication()
	app.AddController(&user.Controller{})

	log.Printf("Starting a server at port %d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), app))
}

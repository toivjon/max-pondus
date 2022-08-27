package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/toivjon/max-pondus/backend/internal/server/admin"
	"github.com/toivjon/max-pondus/backend/internal/server/common/middleware"
	"github.com/toivjon/max-pondus/backend/internal/server/personal"
)

const defaultPort = 8080

func main() {
	port := flag.Int("port", defaultPort, "The port to listen for the incoming HTTP connections.")
	timeout := flag.Duration("timeout", time.Second, "The timeout for processing the request.")
	flag.Parse()

	var personalHandler http.Handler
	personalHandler = personal.NewHandler()
	personalHandler = middleware.BasicAuth("personal", &personal.Authenticator{}, personalHandler)

	var adminHandler http.Handler
	adminHandler = admin.NewHandler()
	adminHandler = middleware.BasicAuth("admin", &admin.Authenticator{}, adminHandler)

	mux := http.NewServeMux()
	mux.Handle("/api/v1/personal", personalHandler)
	mux.Handle("/api/v1/admin", adminHandler)
	mux.Handle("/", http.NotFoundHandler())

	handler := http.TimeoutHandler(mux, *timeout, "")
	handler = middleware.Recoverer(log.Printf, handler)
	handler = middleware.Logger(log.Printf, handler)
	handler = middleware.RequestID(handler)

	server := &http.Server{
		// Just use the default hostname and only specify the port we want to listen.
		Addr: fmt.Sprintf(":%d", *port),
		// Make sure that our response writer has time to write the response when processing timeouts.
		WriteTimeout: *timeout + time.Second,
		// Reserve one second to read the request payload.
		ReadTimeout: time.Second,
		// Reserve one second to request to establish a connection and us to read the headers.
		ReadHeaderTimeout: time.Second,
		// Assign the root handler along with the middleware chain.
		Handler: handler,
		// Tell server to use ReadTimeout value as the timeout value for keep-alive connections.
		IdleTimeout: 0,
		// Use the default maximum amount of bytes for headers.
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
		// Explicitly specify TLS to nil to keep linter happy.
		TLSConfig: nil,
		// Explicitly specify TLS to nil to keep linter happy.
		TLSNextProto: nil,
		// Explicitly specify that we do not need an additional observer for the connection state changes.
		ConnState: nil,
		// Specify server to use standard logger for connection errors.
		ErrorLog: nil,
		// Specify server to use the default context as the base context for all requests.
		BaseContext: nil,
		// Tell server to use the default behaviour for the connection context.
		ConnContext: nil,
	}
	log.Printf("Starting a server port: %d timeout: %v", *port, *timeout)
	log.Fatal(server.ListenAndServe())
}

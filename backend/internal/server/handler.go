package server

import "net/http"

// Handler is the base handler for all HTTP requests.
type Handler struct{}

func (h *Handler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// TODO Wrap response writer and request into a context.
	rw.Write([]byte("TODO"))
}

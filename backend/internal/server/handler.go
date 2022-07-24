package server

import "net/http"

// Handler is the base handler for all HTTP requests.
type Handler struct{}

func (h *Handler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	ctx := Context{ResponseWriter: rw, Request: req}
	ctx.WriteResponse(200, struct{ Text string }{Text: "TODO"})
}

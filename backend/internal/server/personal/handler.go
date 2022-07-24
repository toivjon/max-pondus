package personal

import (
	"net/http"

	"github.com/toivjon/max-pondus/backend/internal/server"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	ctx := server.Context{ResponseWriter: rw, Request: req}
	ctx.WriteResponse(200, struct{ Text string }{Text: "TODO personal!"})
}

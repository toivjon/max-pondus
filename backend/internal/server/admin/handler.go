package admin

import (
	"fmt"
	"net/http"

	"github.com/toivjon/max-pondus/backend/internal/server"
	"github.com/toivjon/max-pondus/backend/internal/server/common/contextkey"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	ctx := server.Context{ResponseWriter: rw, Request: req}
	reqCtx := ctx.Request.Context()
	user := reqCtx.Value(contextkey.User)
	responseText := fmt.Sprintf("TODO Admin: Hello %+v", user)
	ctx.WriteResponse(200, struct{ Text string }{Text: responseText})
}

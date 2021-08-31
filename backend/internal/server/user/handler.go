package user

import (
	"fmt"
	"net/http"
)

// Handler represents a service handling user API requests.
type Handler struct{}

// ServeHTTP will handle requests passed to user API.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello users!")
}

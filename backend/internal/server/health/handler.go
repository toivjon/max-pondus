package health

import (
	"fmt"
	"net/http"
)

// Handler represents a service handling health status requests.
type Handler struct{}

// ServeHTTP will handle the request passed to health API.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello health!")
}

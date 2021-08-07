package muscle

import (
	"fmt"
	"net/http"
)

// Handler represents a service handling muscle API requests.
type Handler struct{}

// ServeHTTP will handle the request passed to muscle API.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello muscles!")
}

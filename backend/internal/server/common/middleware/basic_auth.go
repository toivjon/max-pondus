package middleware

import (
	"fmt"
	"net/http"
)

// TODO Perhaps we should return an User object from the authenticator and add it to context?

// Authenticator is used by the basic auth to check the provided credentials.
type Authenticator interface {
	// Authenticate performs a check whether the given credentials are valid.
	Authenticate(username, password string) bool
}

// BasicAuth performs HTTP Basic Authentication and responds with 401 on invalid credentials.
func BasicAuth(realm string, authenticator Authenticator, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if ok && authenticator.Authenticate(username, password) {
			next.ServeHTTP(w, r)
			return
		}
		headerVal := fmt.Sprintf("Basic realm=%q, charset=%q", realm, "UTF-8")
		w.Header().Set("WWW-Authenticate", headerVal)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}

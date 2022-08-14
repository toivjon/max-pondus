package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/toivjon/max-pondus/backend/internal/server/common"
	"github.com/toivjon/max-pondus/backend/internal/server/common/contextkey"
)

// Authenticator is used by the basic auth to check the provided credentials.
type Authenticator interface {
	// Authenticate performs a check whether the given credentials are valid.
	Authenticate(username, password string) (bool, common.User)
}

// BasicAuth performs HTTP Basic Authentication and responds with 401 on invalid credentials.
func BasicAuth(realm string, authenticator Authenticator, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		authenticated, user := authenticator.Authenticate(username, password)
		if ok && authenticated {
			ctx := r.Context()
			ctx = context.WithValue(ctx, contextkey.User, user)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}
		headerVal := fmt.Sprintf("Basic realm=%q, charset=%q", realm, "UTF-8")
		w.Header().Set("WWW-Authenticate", headerVal)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}

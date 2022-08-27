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
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		username, password, hasCredentials := req.BasicAuth()
		if hasCredentials {
			authenticated, user := authenticator.Authenticate(username, password)
			if authenticated {
				ctx := req.Context()
				ctx = context.WithValue(ctx, contextkey.User, user)
				next.ServeHTTP(res, req.WithContext(ctx))
				return
			}
		}
		headerVal := fmt.Sprintf("Basic realm=%q, charset=%q", realm, "UTF-8")
		res.Header().Set("WWW-Authenticate", headerVal)
		http.Error(res, "Unauthorized", http.StatusUnauthorized)
	})
}

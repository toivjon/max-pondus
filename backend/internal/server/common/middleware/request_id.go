package middleware

import (
	"context"
	"net/http"

	"github.com/toivjon/max-pondus/backend/internal/server/common/contextkey"
	"github.com/toivjon/max-pondus/backend/internal/server/common/random"
)

const requestIDLength = 16

// RequestID assigns a unique identifier to each incoming request.
func RequestID(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, contextkey.RequestID, random.String(requestIDLength))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

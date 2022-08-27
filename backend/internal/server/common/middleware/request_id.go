package middleware

import (
	"context"
	"net/http"

	"github.com/toivjon/max-pondus/backend/internal/server/common/contextkey"
	"github.com/toivjon/max-pondus/backend/internal/server/common/random"
)

const RequestIDLength = 16

// RequestID assigns a unique identifier to each incoming request.
func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, contextkey.RequestID, random.String(RequestIDLength))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

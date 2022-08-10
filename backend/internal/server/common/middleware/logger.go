package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/toivjon/max-pondus/backend/internal/server/common/contextkey"
)

// Logger logs the request-response information after the request has been handled.
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapper := &responseWriterWrapper{ResponseWriter: w}
		defer func() {
			duration := time.Since(start)
			reqID := r.Context().Value(contextkey.RequestID)
			log.Printf("%s %s %s - %d in %s", reqID, r.Method, r.URL, wrapper.statusCode, duration)
		}()
		next.ServeHTTP(wrapper, r)
	})
}

// responseWriterWrapper is a wrapper to capture response data for logging purposes.
type responseWriterWrapper struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader wraps the underlying the WriteHeader call and captures the provided HTTP status.
func (r *responseWriterWrapper) WriteHeader(statusCode int) {
	r.statusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

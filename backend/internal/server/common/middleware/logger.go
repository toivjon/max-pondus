package middleware

import (
	"net/http"
	"time"

	"github.com/toivjon/max-pondus/backend/internal/server/common/contextkey"
)

// LoggerPrintf is used to write a new log entry into the application log.
type LoggerPrintf func(string, ...any)

// Logger logs the request-response information after the request has been handled.
func Logger(loggerPrintf LoggerPrintf, next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		start := time.Now()
		wrapper := &responseWriterWrapper{ResponseWriter: res, statusCode: 0}
		defer func() {
			duration := time.Since(start)
			reqID := req.Context().Value(contextkey.RequestID)
			loggerPrintf("%s %s %s - %d in %s", reqID, req.Method, req.URL, wrapper.statusCode, duration)
		}()
		next.ServeHTTP(wrapper, req)
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

package middleware_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/toivjon/max-pondus/backend/internal/server/common/assert"
	"github.com/toivjon/max-pondus/backend/internal/server/common/contextkey"
	"github.com/toivjon/max-pondus/backend/internal/server/common/middleware"
)

func TestRequestIDAddsRequestIDToContext(t *testing.T) {
	t.Parallel()
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		reqID := getRequestID(ctx)
		assert.Equal(t, middleware.RequestIDLength, len(reqID))
	})
	testHandler(t.Context(), middleware.RequestID(nextHandler))
}

func TestRequestIDOverridesOldRequestID(t *testing.T) {
	t.Parallel()
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		reqID := getRequestID(ctx)
		assert.Equal(t, middleware.RequestIDLength, len(reqID))
	})
	ctx := t.Context()
	ctx = context.WithValue(ctx, contextkey.RequestID, mockVal)
	testHandler(ctx, middleware.RequestID(nextHandler))
}

func TestRequestIDKeepsOldNonRelatedContent(t *testing.T) {
	t.Parallel()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		reqID := getRequestID(ctx)
		otherVal, ok := ctx.Value(mockKey).(string)
		if !ok {
			panic("mockKey is not a string!")
		}
		assert.Equal(t, middleware.RequestIDLength, len(reqID))
		assert.Equal(t, mockVal, otherVal)
	})
	ctx := t.Context()
	ctx = context.WithValue(ctx, mockKey, mockVal)
	testHandler(ctx, middleware.RequestID(handler))
}

func TestRequestIDIsDifferentOnConsecutiveCalls(t *testing.T) {
	t.Parallel()
	reqIDs := make(map[string]bool)
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		reqID := getRequestID(ctx)
		reqIDs[reqID] = true
	})
	handler := middleware.RequestID(nextHandler)
	ctx := t.Context()
	const callCount = 2
	for range callCount {
		testHandler(ctx, handler)
	}
	assert.Equal(t, callCount, len(reqIDs))
}

const (
	mockVal = "mockVal"
	mockKey = contextkey.ContextKey("mockKey")
)

func getRequestID(ctx context.Context) string {
	reqID, ok := ctx.Value(contextkey.RequestID).(string)
	if !ok {
		panic("Context RequestID is not a string!")
	}
	return reqID
}

func testHandler(ctx context.Context, handler http.Handler) {
	req := httptest.NewRequest(http.MethodGet, "http://testing", nil)
	handler.ServeHTTP(httptest.NewRecorder(), req.WithContext(ctx))
}

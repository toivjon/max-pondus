package middleware

import (
	"context"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/toivjon/max-pondus/backend/internal/server/common/assert"
	"github.com/toivjon/max-pondus/backend/internal/server/common/contextkey"
	"github.com/toivjon/max-pondus/backend/internal/server/common/random"
)

func TestRequestIDAddsRequestIDToContext(t *testing.T) {
	rand.Seed(0)
	expected := random.String(requestIDLength)
	rand.Seed(0)

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		assert.Equal(t, expected, ctx.Value(contextkey.RequestID))
	})
	testHandler(context.Background(), RequestID(nextHandler))
}

func TestRequestIDOverridesOldRequestID(t *testing.T) {
	rand.Seed(0)
	expected := random.String(requestIDLength)
	rand.Seed(0)

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		assert.Equal(t, expected, ctx.Value(contextkey.RequestID))
	})
	ctx := context.Background()
	ctx = context.WithValue(ctx, contextkey.RequestID, mockVal)
	testHandler(ctx, RequestID(nextHandler))
}

func TestRequestIDKeepsOldNonRelatedContent(t *testing.T) {
	rand.Seed(0)
	expected := random.String(requestIDLength)
	rand.Seed(0)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		assert.Equal(t, expected, ctx.Value(contextkey.RequestID))
		assert.Equal(t, mockVal, ctx.Value(mockKey))
	})
	ctx := context.Background()
	ctx = context.WithValue(ctx, mockKey, mockVal)
	testHandler(ctx, RequestID(handler))
}

func TestRequestIDIsDifferentOnConsecutiveCalls(t *testing.T) {
	rand.Seed(0)
	expected1 := random.String(requestIDLength)
	expected2 := random.String(requestIDLength)
	rand.Seed(0)

	calls := 0
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		if calls == 0 {
			assert.Equal(t, expected1, ctx.Value(contextkey.RequestID))
		} else {
			assert.Equal(t, expected2, ctx.Value(contextkey.RequestID))
		}
		calls++
	})
	handler := RequestID(nextHandler)
	ctx := context.Background()
	testHandler(ctx, handler)
	testHandler(ctx, handler)
	assert.Equal(t, 2, calls)
}

const mockVal = "mockVal"
const mockKey = contextkey.ContextKey("mockKey")

func testHandler(ctx context.Context, handler http.Handler) {
	req := httptest.NewRequest(http.MethodGet, "http://testing", nil)
	handler.ServeHTTP(httptest.NewRecorder(), req.WithContext(ctx))
}

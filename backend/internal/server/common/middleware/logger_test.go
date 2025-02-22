package middleware_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/toivjon/max-pondus/backend/internal/server/common/assert"
	"github.com/toivjon/max-pondus/backend/internal/server/common/contextkey"
	"github.com/toivjon/max-pondus/backend/internal/server/common/middleware"
)

func TestLoggerWithNoCtxRequestID(t *testing.T) {
	t.Parallel()
	result := ""
	nextHandler := &mockHandler{StatusCode: http.StatusOK, CallCount: 0}
	handler := middleware.Logger(func(format string, args ...any) {
		result = fmt.Sprintf(format, args...)
	}, nextHandler)

	req := httptest.NewRequest(http.MethodGet, "http://testing", nil)
	ctx := t.Context()
	handler.ServeHTTP(httptest.NewRecorder(), req.WithContext(ctx))

	assert.Equal(t, 1, nextHandler.CallCount)
	assert.Match(t, fmt.Sprintf("%s %s %s - [0-9]+ in [0-9]+", "\\(<nil>\\)", req.Method, "http://testing"), result)
}

func TestLoggerWithRequestID(t *testing.T) {
	t.Parallel()
	result := ""
	nextHandler := &mockHandler{StatusCode: http.StatusOK, CallCount: 0}
	handler := middleware.Logger(func(format string, args ...any) {
		result = fmt.Sprintf(format, args...)
	}, nextHandler)

	req := httptest.NewRequest(http.MethodGet, "http://testing", nil)
	ctx := t.Context()
	ctx = context.WithValue(ctx, contextkey.RequestID, mockRequestID)
	handler.ServeHTTP(httptest.NewRecorder(), req.WithContext(ctx))

	assert.Equal(t, 1, nextHandler.CallCount)
	assert.Match(t, fmt.Sprintf("%s %s %s - [0-9]+ in [0-9]+", mockRequestID, req.Method, "http://testing"), result)
}

const mockRequestID = "mockRequestID"

type mockHandler struct {
	StatusCode int
	CallCount  int
}

func (m *mockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.CallCount++
	w.WriteHeader(m.StatusCode)
}

package middleware

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/toivjon/max-pondus/backend/internal/server/common/assert"
	"github.com/toivjon/max-pondus/backend/internal/server/common/contextkey"
)

func TestLoggerWithNoCtxRequestID(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(os.Stderr)
	nextHandler := &mockHandler{StatusCode: http.StatusOK}
	handler := Logger(nextHandler)

	req := httptest.NewRequest(http.MethodGet, "http://testing", nil)
	ctx := context.Background()
	handler.ServeHTTP(httptest.NewRecorder(), req.WithContext(ctx))

	assert.Equal(t, 1, nextHandler.CallCount)
	assert.Match(t, fmt.Sprintf("%s %s %s - [0-9]+ in [0-9]+", "\\(<nil>\\)", http.MethodGet, "http://testing"), buf.String())
}

func TestLoggerWithRequestID(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(os.Stderr)
	nextHandler := &mockHandler{StatusCode: http.StatusOK}
	handler := Logger(nextHandler)

	req := httptest.NewRequest(http.MethodGet, "http://testing", nil)
	ctx := context.Background()
	ctx = context.WithValue(ctx, contextkey.RequestID, mockRequestID)
	handler.ServeHTTP(httptest.NewRecorder(), req.WithContext(ctx))

	assert.Equal(t, 1, nextHandler.CallCount)
	assert.Match(t, fmt.Sprintf("%s %s %s - [0-9]+ in [0-9]+", mockRequestID, http.MethodGet, "http://testing"), buf.String())
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

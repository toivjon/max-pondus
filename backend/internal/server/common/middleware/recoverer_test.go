package middleware_test

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/toivjon/max-pondus/backend/internal/server/common/assert"
	"github.com/toivjon/max-pondus/backend/internal/server/common/contextkey"
	"github.com/toivjon/max-pondus/backend/internal/server/common/middleware"
)

//nolint:paralleltest
func TestRecovererPanicsWithErrAbortHandler(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(os.Stderr)
	nextHandler := &mockPanicHandler{Err: http.ErrAbortHandler, CallCount: 0}
	handler := middleware.Recoverer(nextHandler)
	defer func() {
		panicCount := 0
		if rvr := recover(); rvr != nil {
			panicCount++
		}
		assert.Equal(t, 1, nextHandler.CallCount)
		assert.Equal(t, 1, panicCount)
		assert.Equal(t, "", buf.String())
	}()

	req := httptest.NewRequest(http.MethodGet, "http://testing", nil)
	ctx := context.Background()
	handler.ServeHTTP(httptest.NewRecorder(), req.WithContext(ctx))

	t.Error("Should panic and therefore should not reach this line!")
}

//nolint:paralleltest
func TestRecovererRecoversFromNonErrAbortHandler(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(os.Stderr)
	nextHandler := &mockPanicHandler{Err: errMock, CallCount: 0}
	handler := middleware.Recoverer(nextHandler)

	req := httptest.NewRequest(http.MethodGet, "http://testing", nil)
	ctx := context.Background()
	ctx = context.WithValue(ctx, contextkey.RequestID, mockRequestID)
	handler.ServeHTTP(httptest.NewRecorder(), req.WithContext(ctx))

	assert.Equal(t, 1, nextHandler.CallCount)
	assert.Match(t, fmt.Sprintf("%s panic: %s", mockRequestID, errMock), buf.String())
}

var errMock = errors.New("mockError")

type mockPanicHandler struct {
	Err       error
	CallCount int
}

func (m *mockPanicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.CallCount++
	panic(m.Err)
}

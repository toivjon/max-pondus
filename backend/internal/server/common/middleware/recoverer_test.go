package middleware_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/toivjon/max-pondus/backend/internal/server/common/assert"
	"github.com/toivjon/max-pondus/backend/internal/server/common/contextkey"
	"github.com/toivjon/max-pondus/backend/internal/server/common/middleware"
)

func TestRecovererPanicsWithErrAbortHandler(t *testing.T) {
	t.Parallel()
	result := ""
	nextHandler := &mockPanicHandler{Err: http.ErrAbortHandler, CallCount: 0}
	handler := middleware.Recoverer(func(format string, args ...any) {
		result = fmt.Sprintf(format, args...)
	}, nextHandler)
	defer func() {
		panicCount := 0
		if rvr := recover(); rvr != nil {
			panicCount++
		}
		assert.Equal(t, 1, nextHandler.CallCount)
		assert.Equal(t, 1, panicCount)
		assert.Equal(t, "", result)
	}()

	req := httptest.NewRequest(http.MethodGet, "http://testing", nil)
	ctx := t.Context()
	handler.ServeHTTP(httptest.NewRecorder(), req.WithContext(ctx))

	t.Error("Should panic and therefore should not reach this line!")
}

func TestRecovererRecoversFromNonErrAbortHandler(t *testing.T) {
	t.Parallel()
	result := ""
	nextHandler := &mockPanicHandler{Err: errMock, CallCount: 0}
	handler := middleware.Recoverer(func(format string, args ...any) {
		result = fmt.Sprintf(format, args...)
	}, nextHandler)

	req := httptest.NewRequest(http.MethodGet, "http://testing", nil)
	ctx := t.Context()
	ctx = context.WithValue(ctx, contextkey.RequestID, mockRequestID)
	handler.ServeHTTP(httptest.NewRecorder(), req.WithContext(ctx))

	assert.Equal(t, 1, nextHandler.CallCount)
	assert.Match(t, fmt.Sprintf("%s panic: %s", mockRequestID, errMock), result)
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

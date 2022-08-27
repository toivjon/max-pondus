package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/toivjon/max-pondus/backend/internal/server/common"
	"github.com/toivjon/max-pondus/backend/internal/server/common/assert"
	"github.com/toivjon/max-pondus/backend/internal/server/common/middleware"
)

func TestBasicAuthWithoutBasicAuthDefinition(t *testing.T) {
	t.Parallel()
	authenticator := &mockAuthenticator{CallCount: 0, Result: false, ResultUser: common.User{Username: ""}}
	nextHandler := new(mockHandler)
	handler := middleware.BasicAuth(mockRealm, authenticator, nextHandler)

	req := httptest.NewRequest(http.MethodGet, "http://testing", nil)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	assert.Equal(t, 0, authenticator.CallCount)
	assert.Equal(t, 0, nextHandler.CallCount)
	assert.Equal(t, http.StatusUnauthorized, rr.Code)
}

func TestBasicAuthFailure(t *testing.T) {
	t.Parallel()
	authenticator := &mockAuthenticator{CallCount: 0, Result: false, ResultUser: common.User{Username: ""}}
	nextHandler := new(mockHandler)
	handler := middleware.BasicAuth(mockRealm, authenticator, nextHandler)

	req := httptest.NewRequest(http.MethodGet, "http://testing", nil)
	req.SetBasicAuth(mockUsername, mockPassword)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	assert.Equal(t, 1, authenticator.CallCount)
	assert.Equal(t, 0, nextHandler.CallCount)
	assert.Equal(t, http.StatusUnauthorized, rr.Code)
}

func TestBasicAuthSuccess(t *testing.T) {
	t.Parallel()
	authenticator := &mockAuthenticator{CallCount: 0, Result: true, ResultUser: common.User{Username: ""}}
	nextHandler := &mockHandler{StatusCode: http.StatusOK, CallCount: 0}
	handler := middleware.BasicAuth(mockRealm, authenticator, nextHandler)

	req := httptest.NewRequest(http.MethodGet, "http://testing", nil)
	req.SetBasicAuth(mockUsername, mockPassword)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	assert.Equal(t, 1, authenticator.CallCount)
	assert.Equal(t, 1, nextHandler.CallCount)
	assert.Equal(t, http.StatusOK, rr.Code)
}

const (
	mockRealm    = "mockRealm"
	mockUsername = "mockUsername"
	mockPassword = "mockPassword"
)

type mockAuthenticator struct {
	CallCount  int
	Result     bool
	ResultUser common.User
}

func (m *mockAuthenticator) Authenticate(username, password string) (bool, common.User) {
	m.CallCount++
	return m.Result, m.ResultUser
}

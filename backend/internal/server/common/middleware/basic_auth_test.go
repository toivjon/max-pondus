package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/toivjon/max-pondus/backend/internal/server/common"
	"github.com/toivjon/max-pondus/backend/internal/server/common/assert"
)

func TestBasicAuthWithoutBasicAuthDefinition(t *testing.T) {
	authenticator := &mockAuthenticator{Result: false}
	nextHandler := new(mockHandler)
	handler := BasicAuth(mockRealm, authenticator, nextHandler)

	req := httptest.NewRequest(http.MethodGet, "http://testing", nil)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	assert.Equal(t, 0, authenticator.CallCount)
	assert.Equal(t, 0, nextHandler.CallCount)
	assert.Equal(t, http.StatusUnauthorized, rr.Code)
}

func TestBasicAuthFailure(t *testing.T) {
	authenticator := &mockAuthenticator{Result: false}
	nextHandler := new(mockHandler)
	handler := BasicAuth(mockRealm, authenticator, nextHandler)

	req := httptest.NewRequest(http.MethodGet, "http://testing", nil)
	req.SetBasicAuth(mockUsername, mockPassword)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	assert.Equal(t, 1, authenticator.CallCount)
	assert.Equal(t, 0, nextHandler.CallCount)
	assert.Equal(t, http.StatusUnauthorized, rr.Code)
}

func TestBasicAuthSuccess(t *testing.T) {
	authenticator := &mockAuthenticator{Result: true}
	nextHandler := &mockHandler{StatusCode: http.StatusOK}
	handler := BasicAuth(mockRealm, authenticator, nextHandler)

	req := httptest.NewRequest(http.MethodGet, "http://testing", nil)
	req.SetBasicAuth(mockUsername, mockPassword)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	assert.Equal(t, 1, authenticator.CallCount)
	assert.Equal(t, 1, nextHandler.CallCount)
	assert.Equal(t, http.StatusOK, rr.Code)
}

const mockRealm = "mockRealm"
const mockUsername = "mockUsername"
const mockPassword = "mockPassword"

type mockAuthenticator struct {
	CallCount  int
	Result     bool
	ResultUser common.User
}

func (m *mockAuthenticator) Authenticate(username, password string) (bool, common.User) {
	m.CallCount++
	return m.Result, m.ResultUser
}

package server_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/toivjon/max-pondus/backend/internal/server"
)

func TestContextWriteResponse(t *testing.T) {
	t.Parallel()
	res := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	ctx := &server.Context{ResponseWriter: res, Request: req}

	dto := struct{ Val string }{Val: "foo"}
	expectedCode := http.StatusOK
	ctx.WriteResponse(expectedCode, dto)

	assertCode(t, res, expectedCode)
	assertContentType(t, res, "application/json")
	assertBody(t, res, "{\"Val\":\"foo\"}\n")
}

func assertCode(t *testing.T, res *httptest.ResponseRecorder, expectedCode int) {
	t.Helper()
	if res.Code != expectedCode {
		t.Fatalf("Unexpected status code: %d, expected: %d", res.Code, expectedCode)
	}
}

func assertContentType(t *testing.T, rw *httptest.ResponseRecorder, expectedType string) {
	t.Helper()
	contentType := rw.Header().Get("Content-Type")
	if contentType != expectedType {
		t.Fatalf("Unexpected content type: %s, expected: %s", contentType, expectedType)
	}
}

func assertBody(t *testing.T, rw *httptest.ResponseRecorder, expectedBody string) {
	t.Helper()
	expectedBytes := bytes.NewBufferString(expectedBody)
	if !bytes.Equal(rw.Body.Bytes(), expectedBytes.Bytes()) {
		t.Fatalf("Unexpected body: %q, expected: %q", rw.Body, expectedBytes)
	}
}

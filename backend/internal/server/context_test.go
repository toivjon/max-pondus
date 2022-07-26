package server

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestContextWriteResponse(t *testing.T) {
	rw := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	ctx := &Context{ResponseWriter: rw, Request: req}

	dto := struct{ Val string }{Val: "foo"}
	expectedCode := http.StatusOK
	ctx.WriteResponse(expectedCode, dto)

	assertCode(t, rw, expectedCode)
	assertContentType(t, rw, "application/json")
	assertBody(t, rw, "{\"Val\":\"foo\"}\n")
}

func assertCode(t *testing.T, rw *httptest.ResponseRecorder, expectedCode int) {
	t.Helper()
	if rw.Code != expectedCode {
		t.Fatalf("Unexpected status code: %d, expected: %d", rw.Code, expectedCode)
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

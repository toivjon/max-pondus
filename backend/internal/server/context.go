package server

import (
	"encoding/json"
	"net/http"
)

// Context contains the process context and wraps the original request reader and response writer.
type Context struct {
	http.ResponseWriter
	*http.Request
}

// WriteResponse can be used to write the HTTP status and JSON payload as the response.
func (c *Context) WriteResponse(statusCode int, dto interface{}) {
	c.ResponseWriter.Header().Set("Content-Type", "application/json")
	c.WriteHeader(statusCode)
	json.NewEncoder(c.ResponseWriter).Encode(dto)
}

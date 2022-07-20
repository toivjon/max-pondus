package server

import "regexp"

// Route presents a single request route which has a URL pattern, method and request handler.
type Route struct {
	pattern *regexp.Regexp
	method  string
	handler Handler
}

// NewRoute builds a new Route for the given URL pattern, method and request handler.
func NewRoute(pattern, method string, handler Handler) Route {
	re := regexp.MustCompile(pattern)
	return Route{pattern: re, handler: handler, method: method}
}

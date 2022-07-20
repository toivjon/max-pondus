package server

// Handler represents a function which can handle an incoming HTTP request.
type Handler func(*Context)

package server

// Controller handles incoming API requests.
type Controller interface {
	// Routes returns a list of end points provided by the controller.
	Routes() []Route
}

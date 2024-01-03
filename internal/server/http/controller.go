package http

import "net/http"

/* __________________________________________________ */

// Controller ...
type Controller interface {
	Handle(writer http.ResponseWriter, request *http.Request)
}

// AsHandler ...
func AsHandler(c Controller) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		c.Handle(writer, request)
	}
}

/* __________________________________________________ */

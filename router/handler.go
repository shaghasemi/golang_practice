package router

import "net/http"

// Handler is the contract our mini framework expects.
// Anything with a ServeHTTP method can be registered as a route handler.
type Handler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

// HandlerFunc turns a plain function into a Handler.
//
// This is the adapter pattern: you write func(w, r), the framework wants
// an interface. HandlerFunc bridges the two.
type HandlerFunc func(http.ResponseWriter, *http.Request)

func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}

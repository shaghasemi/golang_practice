package router

import (
	"net/http"
	"strings"
)

type Router struct {
	routes map[string]Handler
}

func New() *Router {
	return &Router{routes: make(map[string]Handler)}
}

// Handle registers any Handler — struct, mock, middleware wrapper, or HandlerFunc.
func (rt *Router) Handle(path string, h Handler) {
	rt.routes[path] = h
}

// HandleFunc is sugar for registering a plain function.
// Internally it converts the function to HandlerFunc, which satisfies Handler.
func (rt *Router) HandleFunc(path string, fn func(http.ResponseWriter, *http.Request)) {
	rt.Handle(path, HandlerFunc(fn))
}

func (rt *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSuffix(r.URL.Path, "/")
	if path == "" {
		path = "/"
	}

	h, ok := rt.routes[path]
	if !ok {
		http.NotFound(w, r)
		return
	}

	h.ServeHTTP(w, r)
}

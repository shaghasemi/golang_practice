package router

import (
	"net/http"
	"strings"
)

// LegacyRouter mirrors Router but uses the struct-wrapper adapter instead of HandlerFunc.
// Both routers expose the same API; only HandleFunc's internal conversion differs.
type LegacyRouter struct {
	routes map[string]Handler
}

func NewLegacy() *LegacyRouter {
	return &LegacyRouter{routes: make(map[string]Handler)}
}

func (rt *LegacyRouter) Handle(path string, h Handler) {
	rt.routes[path] = h
}

func (rt *LegacyRouter) HandleFunc(path string, fn func(http.ResponseWriter, *http.Request)) {
	rt.Handle(path, LegacyFromFunc(fn))
}

func (rt *LegacyRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

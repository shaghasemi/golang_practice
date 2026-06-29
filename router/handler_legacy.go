package router

import "net/http"

// funcHandler is the pre-function-type way to turn a plain func into a Handler.
// Before Go let you write `type HandlerFunc func(...)` with methods, frameworks
// invented a small struct that holds the function and delegates ServeHTTP to it.
type funcHandler struct {
	fn func(http.ResponseWriter, *http.Request)
}

func (h funcHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.fn(w, r)
}

// LegacyFromFunc wraps fn in a struct so it satisfies Handler.
// Same job as HandlerFunc(fn), different mechanism.
func LegacyFromFunc(fn func(http.ResponseWriter, *http.Request)) Handler {
	return funcHandler{fn: fn}
}

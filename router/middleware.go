package router

import "net/http"

// MiddlewareFunc is another function type with methods.
// It wraps one Handler and returns another — the classic middleware shape.
type MiddlewareFunc func(Handler) Handler

func (mw MiddlewareFunc) Then(next Handler) Handler {
	return mw(next)
}

// Chain applies middleware left-to-right: Chain(a, b)(handler) => a(b(handler)).
func Chain(middlewares ...MiddlewareFunc) MiddlewareFunc {
	return MiddlewareFunc(func(final Handler) Handler {
		h := final
		for i := len(middlewares) - 1; i >= 0; i-- {
			h = middlewares[i](h)
		}
		return h
	})
}

func Logging() MiddlewareFunc {
	return MiddlewareFunc(func(next Handler) Handler {
		return HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// In real code you'd log method, path, duration, etc.
			next.ServeHTTP(w, r)
		})
	})
}

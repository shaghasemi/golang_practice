package router

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// --- handlers used in tests ---

type greetHandler struct {
	message string
}

func (h greetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(h.message))
}

func writeStatusOK(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// --- tests that show why HandlerFunc exists ---

func TestPlainFunctionDoesNotSatisfyHandler(t *testing.T) {
	var _ Handler = greetHandler{message: "hi"} // struct works

	// Uncommenting the next line would not compile:
	//
	//   var _ Handler = writeStatusOK
	//
	// writeStatusOK is a func value, not a type with methods.
	// Go interfaces are satisfied by methods on the value's type.
	_ = writeStatusOK
}

func TestHandlerFuncSatisfiesHandler(t *testing.T) {
	// Conversion gives the function a named type, which carries the ServeHTTP method.
	var _ Handler = HandlerFunc(writeStatusOK)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	HandlerFunc(writeStatusOK).ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}
}

func TestRouterAcceptsBothStructAndFunctionHandlers(t *testing.T) {
	rt := New()
	rt.Handle("/struct", greetHandler{message: "from struct"})
	rt.HandleFunc("/func", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("from func"))
	})

	t.Run("struct handler", func(t *testing.T) {
		rec := httptest.NewRecorder()
		rt.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/struct", nil))

		if body := rec.Body.String(); body != "from struct" {
			t.Fatalf("body = %q, want %q", body, "from struct")
		}
	})

	t.Run("function handler", func(t *testing.T) {
		rec := httptest.NewRecorder()
		rt.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/func", nil))

		if body := rec.Body.String(); body != "from func" {
			t.Fatalf("body = %q, want %q", body, "from func")
		}
	})
}

func TestMiddlewareFuncWrapsHandlers(t *testing.T) {
	called := false

	base := HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		called = true
		w.WriteHeader(http.StatusOK)
	})

	wrapped := Logging().Then(base)

	rec := httptest.NewRecorder()
	wrapped.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/", nil))

	if !called {
		t.Fatal("expected base handler to run through middleware")
	}
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}
}

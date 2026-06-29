package router

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Both adapters turn the same plain function into a Handler.
func TestModernAndLegacyAdaptersBothSatisfyHandler(t *testing.T) {
	fn := func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}

	var modern Handler = HandlerFunc(fn)
	var legacy Handler = LegacyFromFunc(fn)

	for name, h := range map[string]Handler{
		"HandlerFunc":  modern,
		"funcHandler": legacy,
	} {
		t.Run(name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			h.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/", nil))

			if rec.Code != http.StatusOK {
				t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
			}
		})
	}
}

func TestModernAndLegacyRoutersBehaveTheSame(t *testing.T) {
	handler := func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("hello"))
	}

	cases := []struct {
		name   string
		router http.Handler
	}{
		{
			name: "modern HandlerFunc router",
			router: func() http.Handler {
				rt := New()
				rt.HandleFunc("/", handler)
				return rt
			}(),
		},
		{
			name: "legacy funcHandler router",
			router: func() http.Handler {
				rt := NewLegacy()
				rt.HandleFunc("/", handler)
				return rt
			}(),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			tc.router.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/", nil))

			if body := rec.Body.String(); body != "hello" {
				t.Fatalf("body = %q, want %q", body, "hello")
			}
		})
	}
}

func TestBothAdaptersWorkWithMiddleware(t *testing.T) {
	called := false
	fn := func(w http.ResponseWriter, _ *http.Request) {
		called = true
		w.WriteHeader(http.StatusOK)
	}

	for name, base := range map[string]Handler{
		"HandlerFunc":  HandlerFunc(fn),
		"funcHandler": LegacyFromFunc(fn),
	} {
		t.Run(name, func(t *testing.T) {
			called = false
			wrapped := Logging().Then(base)

			rec := httptest.NewRecorder()
			wrapped.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/", nil))

			if !called {
				t.Fatal("expected handler to run through middleware")
			}
		})
	}
}

// BenchmarkRegister compares the cost of wrapping a function at registration time.
// Run with: go test -bench=BenchmarkRegister -benchmem ./router/...
func BenchmarkRegisterModern(b *testing.B) {
	fn := func(http.ResponseWriter, *http.Request) {}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		rt := New()
		rt.HandleFunc("/", fn)
		_ = rt
	}
}

func BenchmarkRegisterLegacy(b *testing.B) {
	fn := func(http.ResponseWriter, *http.Request) {}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		rt := NewLegacy()
		rt.HandleFunc("/", fn)
		_ = rt
	}
}

// BenchmarkAdapter measures only the conversion to Handler (stored in an interface
// so the compiler cannot elide the boxing step).
func BenchmarkAdapterModern(b *testing.B) {
	fn := func(http.ResponseWriter, *http.Request) {}

	b.ReportAllocs()
	var sink Handler
	for i := 0; i < b.N; i++ {
		sink = HandlerFunc(fn)
	}
	_ = sink
}

func BenchmarkAdapterLegacy(b *testing.B) {
	fn := func(http.ResponseWriter, *http.Request) {}

	b.ReportAllocs()
	var sink Handler
	for i := 0; i < b.N; i++ {
		sink = LegacyFromFunc(fn)
	}
	_ = sink
}

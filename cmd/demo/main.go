package main

import (
	"fmt"
	"log"
	"net/http"

	"golang_practice/router"
)

func main() {
	rt := router.New()

	// Register a plain function — no struct boilerplate.
	rt.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, "Hello from a function handler")
	})

	// Register a struct handler when you need state.
	rt.Handle("/version", versionHandler{version: "1.0.0"})

	// Middleware still works because both sides speak Handler.
	rt.Handle("/protected", router.Chain(router.Logging())(
		router.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			fmt.Fprintln(w, "Protected route")
		}),
	))

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", rt))
}

type versionHandler struct {
	version string
}

func (h versionHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "version: %s\n", h.version)
}

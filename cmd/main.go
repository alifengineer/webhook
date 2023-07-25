package main

import (
	"go-back-redirector/cmd/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// Simple rest api
	r := mux.NewRouter()

	// Routes consist of a path and a handler function.
	r.HandleFunc("/api/v1/health", handler.HealthHandler).Methods("GET")

	// Bind to a port and pass our router in
	err := http.ListenAndServe(":5000", r)
	if err != nil {
		panic(err)
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// setupRouter configures and returns the HTTP mux for the server.
func setupRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[REQUEST] %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
		fmt.Fprintf(w, "Hello from GKE! You requested: %s\n", r.URL.Path)
	})

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	return mux
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := setupRouter()

	log.Printf("Server listening on port %s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Server failed: %s", err)
	}
}

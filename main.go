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
		port = "8089"
	}

	router := setupRouter()

	log.Printf("Server listening on port %s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Server failed: %s", err)
	}
}

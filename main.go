package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// getEnv fetches a variable or returns a fallback default
func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

func main() {
	port := getEnv("PORT", "8080")
	appName := getEnv("APP_NAME", "go-server")
	environment := getEnv("ENVIRONMENT", "development")

	log.Printf("Starting %s in %s mode on port :%s", appName, environment, port)

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from %s running in %s!", appName, environment)
	})

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

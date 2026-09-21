package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetEnvFallback(t *testing.T) {
	// 1. Test when the variable is NOT set (fallback used)
	result := getEnv("NON_EXISTENT_KEY", "default_val")
	if result != "default_val" {
		t.Errorf("expected 'default_val', got '%s'", result)
	}

	// 2. Test when the variable IS set
	t.Setenv("TEST_APP_NAME", "test-server")
	result = getEnv("TEST_APP_NAME", "default_val")
	if result != "test-server" {
		t.Errorf("expected 'test-server', got '%s'", result)
	}
}

func TestHealthCheck(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	w := httptest.NewRecorder()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}

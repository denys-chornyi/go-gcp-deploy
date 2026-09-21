package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealthCheckHandler(t *testing.T) {
	router := setupRouter()

	req, err := http.NewRequest(http.MethodGet, "/healthz", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Verify Status Code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Verify Response Body
	expected := "ok"
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestRootHandler(t *testing.T) {
	router := setupRouter()

	tests := []struct {
		name         string
		path         string
		expectedBody string
	}{
		{
			name:         "Root path",
			path:         "/",
			expectedBody: "Hello from GKE! You requested: /\n",
		},
		{
			name:         "Custom subpath",
			path:         "/api/v1/resource",
			expectedBody: "Hello from GKE! You requested: /api/v1/resource\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, tt.path, nil)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusOK {
				t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
			}

			if rr.Body.String() != tt.expectedBody {
				t.Errorf("Handler returned unexpected body: got %q want %q", rr.Body.String(), tt.expectedBody)
			}
		})
	}
}

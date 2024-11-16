package router

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	// Create a request to pass to our handler
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(homeHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body
	expected := "Welcome to the Home Page! 🏠"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestV2HelloHandler(t *testing.T) {
	// Create a request
	req, err := http.NewRequest("GET", "/api/v2/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(v2HelloHandler)
	handler.ServeHTTP(rr, req)

	// Check status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Parse JSON response
	var response map[string]interface{}
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Errorf("Failed to decode JSON response: %v", err)
	}

	// Check response fields
	if msg, ok := response["message"].(string); !ok || msg != "Hello from API v2! 👋" {
		t.Errorf("handler returned unexpected message: got %v", response["message"])
	}

	if version, ok := response["version"].(float64); !ok || version != 2 {
		t.Errorf("handler returned unexpected version: got %v", response["version"])
	}
}

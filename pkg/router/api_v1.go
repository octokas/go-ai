package router

import (
	"encoding/json"
	"net/http"
)

func setupV1Routes() {
	http.HandleFunc("/api/v1/hello", v1HelloHandler)
	http.HandleFunc("/api/v1/users", v1UsersHandler)
	http.HandleFunc("/api/v1/status", v1StatusHandler)
}

func v1HelloHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Hello from API v1! 👋"}
	json.NewEncoder(w).Encode(response)
}

func v1UsersHandler(w http.ResponseWriter, r *http.Request) {
	users := []map[string]string{
		{"id": "1", "name": "Alice"},
		{"id": "2", "name": "Bob"},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func v1StatusHandler(w http.ResponseWriter, r *http.Request) {
	status := map[string]string{
		"status":  "healthy",
		"version": "1.0.0",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}

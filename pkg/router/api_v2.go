package router

import (
	"encoding/json"
	"net/http"
)

func setupV2Routes() {
	http.HandleFunc("/api/v2/hello", v2HelloHandler)
	http.HandleFunc("/api/v2/users", v2UsersHandler)
	http.HandleFunc("/api/v2/status", v2StatusHandler)
}

func v2HelloHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"message": "Hello from API v2! 👋",
		"version": 2,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func v2UsersHandler(w http.ResponseWriter, r *http.Request) {
	users := []map[string]interface{}{
		{
			"id":     "1",
			"name":   "Alice",
			"role":   "admin",
			"active": true,
		},
		{
			"id":     "2",
			"name":   "Bob",
			"role":   "user",
			"active": true,
		},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func v2StatusHandler(w http.ResponseWriter, r *http.Request) {
	status := map[string]interface{}{
		"status":  "healthy",
		"version": "2.0.0",
		"services": map[string]string{
			"database": "connected",
			"cache":    "connected",
			"queue":    "active",
		},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}

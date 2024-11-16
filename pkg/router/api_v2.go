package router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/octokas/go-ai/pkg/config"
	"github.com/octokas/go-ai/pkg/logger"
	"github.com/octokas/go-ai/pkg/server"
)

func setupV2Routes() {
	http.HandleFunc("/api/v2/hello", v2HelloHandler)
	http.HandleFunc("/api/v2/users", v2UsersHandler)
	http.HandleFunc("/api/v2/status", v2StatusHandler)
}

func v2HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from API v2! 👋")
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

func V2Server() { // Initial logging with standard log package
	log.Println("Initializing application...")

	// Initialize logger
	logger := logger.New()
	logger.Info("Starting APIv2 server...")

	// Setup home routes
	setupV2Routes()

	// Start server
	log.Printf("[INFO] Server on port :3030")
	http.ListenAndServe(":3030", nil)

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		logger.Fatal("Failed to load configuration:", err)
	}

	// Initialize server
	srv := server.New(cfg)

	// Start server
	if err := srv.Start(); err != nil && err != http.ErrServerClosed {
		logger.Fatal("Server failed:", err)
	}
}

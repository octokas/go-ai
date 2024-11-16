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

func setupV1Routes() {
	http.HandleFunc("/api/v1/hello", v1HelloHandler)
	http.HandleFunc("/api/v1/users", v1UsersHandler)
	http.HandleFunc("/api/v1/status", v1StatusHandler)
}

func v1HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from API v1! 👋")
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

func V1Server() { // Initial logging with standard log package
	log.Println("Initializing application...")

	// Initialize logger
	logger := logger.New()
	logger.Info("Starting APIv2 server...")

	// Setup home routes
	setupV1Routes()

	// Start server
	log.Printf("[INFO] Server on port :2020")
	http.ListenAndServe(":2020", nil)

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

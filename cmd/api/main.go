package main

import (
	"log"
	"net/http"

	"github.com/yourusername/projectname/internal/server"
	"github.com/yourusername/projectname/internal/config"
	"github.com/yourusername/projectname/internal/logger"
)

func main() {
	// Initialize logger
	log := logger.New()
	log.Info("Starting API server...")

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	// Initialize server
	srv := server.New(cfg)

	// Start server
	if err := srv.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server failed:", err)
	}
} 
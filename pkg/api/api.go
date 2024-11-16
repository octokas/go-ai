package api

import (
	"log"
	"net/http"

	"github.com/octokas/go-ai/pkg/config"
	"github.com/octokas/go-ai/pkg/logger"
	"github.com/octokas/go-ai/pkg/server"
)

func RunAPI() {
	// Initial logging with standard log package
	log.Println("Initializing application...")

	// Initialize logger
	logger := logger.New()
	logger.Info("Starting API server...")

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

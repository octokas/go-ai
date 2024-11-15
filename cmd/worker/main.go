package main

import (
	"log"

	"github.com/octokas/go-ai/internal/config"
	"github.com/octokas/go-ai/internal/logger"
)

func main() {
	// Initialize logger
	log := logger.New()
	log.Info("Starting worker...")

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	// TODO: Initialize worker pipeline
} 
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"

	"go-kas/config"
	"go-kas/databases"
	"go-kas/logging"
	"go-kas/server"
)

// Global variables for app-wide access
var (
	cfg    *config.Config
	db     *databases.Database
	logger *logging.Logger
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	// Initialize logger
	logger := logging.NewLogger()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	logger.InfoLog.Println("Starting server...")

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		logger.ErrorLog.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize database connections
	db, err := databases.NewDatabase(cfg)
	if err != nil {
		logger.ErrorLog.Fatalf("Failed to connect to database: %v", err)
	}

	// Ping database to verify connection
	if err := db.Ping(); err != nil {
		logger.ErrorLog.Fatalf("Failed to ping database: %v", err)
	}
	// if err := db.MongoDB.Client().Ping(context.Background(), nil); err != nil {
	// 	logger.ErrorLog.Fatalf("Failed to ping database: %v", err)
	// }

	// Initialize server
	srv := server.NewServer(cfg, db, logger)
	//srv := SetupServer()

	// Start server in a goroutine
	go func() {
		logger.InfoLog.Printf("Server starting on port %s", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.ErrorLog.Fatalf("Server failed: %v", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.InfoLog.Println("Server is shutting down...")

	// Create a deadline for server shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := srv.Shutdown(ctx); err != nil {
		logger.ErrorLog.Printf("Server forced to shutdown: %v", err)
	}

	logger.InfoLog.Println("Server exited properly")
}

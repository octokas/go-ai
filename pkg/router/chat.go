package router

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/octokas/go-ai/pkg/chat"
	"github.com/octokas/go-ai/pkg/config"
	"github.com/octokas/go-ai/pkg/logger"
	"github.com/octokas/go-ai/pkg/server"
)

func setupChatRoutes() {
	http.HandleFunc("/chat/prompt", chatPromptHandler)
	http.HandleFunc("/chat/api", chat.HandleAPI)
}

func chatPromptHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/pagerprompt.html")
}

func RunChatServer(service *chat.Service) error {
	r := gin.Default()

	log.Println("Initializing application...")

	// Initialize logger
	logger := logger.New()
	logger.Info("Starting Chat server...")

	// Setup chat routes
	setupChatRoutes()

	// Start server
	log.Printf("[INFO] Chat server starting on port :4040")
	http.ListenAndServe(":4040", nil)

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

	return r.Run(":4040")
}

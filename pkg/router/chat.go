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

func setupChatRoutes(chatHandler *chat.Handler) {
	http.HandleFunc("/chat/prompt", chatPromptHandler)
	http.HandleFunc("/chat/api", chatHandler.HandleAPI)
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

	// Create handler with service
	chatHandler := chat.NewHandler(service)

	// Setup chat routes
	setupChatRoutes(chatHandler)

	// Update this line to use the handler's method
	http.HandleFunc("/chat/api", chatHandler.HandleAPI)

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		logger.Fatal("Failed to load configuration:", err)
	}

	// Initialize server
	srv := server.New(cfg)

	// Start server
	log.Printf("[INFO] Chat server starting on port :4040")
	http.ListenAndServe(":4040", nil)
	if err := srv.Start(); err != nil && err != http.ErrServerClosed {
		logger.Fatal("Server failed:", err)
	}

	return r.Run(":4040")
}

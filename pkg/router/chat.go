package router

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/octokas/go-ai/pkg/chat"
	"github.com/octokas/go-ai/pkg/config"
	"github.com/octokas/go-ai/pkg/logger"
	"github.com/octokas/go-ai/pkg/server"
)

func setupChatRoutes(r *gin.Engine, chatHandler *chat.Handler) {
	r.GET("/chat/prompt", func(c *gin.Context) {
		c.File("templates/pagerprompt.html")
	})
	r.POST("/chat/api", func(c *gin.Context) {
		chatHandler.HandleGinAPI(c)
	})
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
	setupChatRoutes(r, chatHandler)

	// Load configuration
	// cfg, err := config.Load()
	// if err != nil {
	// 	logger.Fatal("Failed to load configuration:", err)
	// 	return err
	// }

	// Initialize server with proper port formatting
	// srv := server.New(&config.Config{
	// 	Server: config.ServerConfig{
	// 		Port:    cfg.Port,
	// 		Host:    "",
	// 		Timeout: 30,
	// 	},
	// })

	const chatPort = 4040

	srv := server.New(&config.Config{
		Server: config.ServerConfig{
			Port:    chatPort,
			Host:    "",
			Timeout: 30,
		},
	})

	// Start server with properly formatted port
	log.Printf("[INFO] Chat server starting on port :%d", chatPort)
	return srv.Run(r)
}

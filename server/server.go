package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"go-kas/config"
	"go-kas/databases"
	"go-kas/logging"
	"go-kas/middleware"
	"go-kas/routers"
)

type Server struct {
	*http.Server
	config *config.Config
	db     *databases.Database
	logger *logging.Logger
}

func NewServer(cfg *config.Config, db *databases.Database, logger *logging.Logger) *Server {
	// Initialize router with dependencies
	router := routers.SetupRoutes()

	// Apply global middleware
	handler := middleware.MiddlewareChain(
		router,
		middleware.MiddlewareLogger(logger.InfoLog),
		middleware.MiddlewareRecovery(logger.ErrorLog),
		middleware.CORS,
	)

	return &Server{
		Server: &http.Server{
			Addr:         fmt.Sprintf(":%s", cfg.Port),
			Handler:      handler,
			ReadTimeout:  15 * time.Second,
			WriteTimeout: 15 * time.Second,
			IdleTimeout:  60 * time.Second,
		},
		config: cfg,
		db:     db,
		logger: logger,
	}
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.Server.Shutdown(ctx)
}

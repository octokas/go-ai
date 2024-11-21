package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"go-kas/config"
	"go-kas/middleware"
	"go-kas/routers"
)

type Server struct {
	*http.Server
}

func NewServer(cfg *config.Config) *Server {
	router := routers.SetupRoutes()

	// Apply global middleware
	handler := middleware.Chain(
		router,
		middleware.Logger,
		//middleware.Recovery,
		middleware.CORS,
	)

	return &Server{
		&http.Server{
			Addr:         fmt.Sprintf(":%s", cfg.Port),
			Handler:      handler,
			ReadTimeout:  15 * time.Second,
			WriteTimeout: 15 * time.Second,
			IdleTimeout:  60 * time.Second,
		},
	}
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.Server.Shutdown(ctx)
}

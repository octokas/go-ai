package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/yourusername/projectname/internal/config"
	"github.com/yourusername/projectname/internal/logger"
)

type Server struct {
	server *http.Server
	logger *logger.Logger
	config *config.Config
}

func New(cfg *config.Config) *Server {
	log := logger.New()

	mux := http.NewServeMux()
	setupRoutes(mux)

	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port),
		Handler:      mux,
		ReadTimeout:  time.Duration(cfg.Server.Timeout) * time.Second,
		WriteTimeout: time.Duration(cfg.Server.Timeout) * time.Second,
	}

	return &Server{
		server: srv,
		logger: log,
		config: cfg,
	}
}

func setupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/health", healthCheckHandler)
	// Add more routes here
}

func (s *Server) Start() error {
	s.logger.Info("Starting server on", s.server.Addr)
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.logger.Info("Shutting down server...")
	return s.server.Shutdown(ctx)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

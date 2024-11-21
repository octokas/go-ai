package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/octokas/go-kas/go/internal"
	"github.com/octokas/go-kas/go/pkg/config"
	"go.uber.org/zap"
)

type Server struct {
	config     *config.Config
	logger     *zap.Logger
	httpServer *http.Server
	app        *internal.AppContext
}

func New() (*Server, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("loading config: %w", err)
	}

	logger, err := newLogger(cfg.Logging)
	if err != nil {
		return nil, fmt.Errorf("creating logger: %w", err)
	}

	app, err := internal.NewAppContext(context.Background(), cfg)
	if err != nil {
		return nil, fmt.Errorf("creating app context: %w", err)
	}

	server := &Server{
		config: cfg,
		logger: logger,
		app:    app,
	}

	return server, nil
}

func (s *Server) Start() error {
	router := NewRouter(s.app)

	s.httpServer = &http.Server{
		Addr:         fmt.Sprintf("%s:%d", s.config.Server.Host, s.config.Server.Port),
		Handler:      router,
		ReadTimeout:  s.config.Server.ReadTimeout,
		WriteTimeout: s.config.Server.WriteTimeout,
	}

	// Channel for graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		s.logger.Info("starting server",
			zap.String("address", s.httpServer.Addr),
			zap.String("environment", s.config.Server.Environment))

		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Fatal("server error", zap.Error(err))
		}
	}()

	<-stop

	return s.Shutdown()
}

func (s *Server) Shutdown() error {
	s.logger.Info("shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("server shutdown: %w", err)
	}

	if err := s.app.Shutdown(ctx); err != nil {
		return fmt.Errorf("app shutdown: %w", err)
	}

	return nil
}

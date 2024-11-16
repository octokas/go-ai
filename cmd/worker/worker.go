package worker

import (
	"github.com/octokas/go-ai/internal/config"
	"github.com/octokas/go-ai/internal/logger"
)

type Worker struct {
	config *config.Config
	logger *logger.Logger
}

func NewWorker(cfg *config.Config, log *logger.Logger) *Worker {
	return &Worker{
		config: cfg,
		logger: log,
	}
}

func (w *Worker) Run() error {
	w.logger.Info("Worker pipeline started")
	// TODO: Add actual work processing here
	return nil
}

func SetupWorker() {
	// Initialize logger
	log := logger.New()
	log.Info("Starting worker...")

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	// Initialize and run worker
	worker := NewWorker(cfg, log)
	if err := worker.Run(); err != nil {
		log.Fatal("Worker pipeline failed:", err)
	}
}

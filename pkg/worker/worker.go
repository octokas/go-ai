package worker

import (
	"github.com/octokas/go-ai/pkg/config"
	"github.com/octokas/go-ai/pkg/logger"
)

type Worker struct {
	config config.Configer
	logger logger.LoggerInterface
}

func NewWorker(config config.Configer, logger logger.LoggerInterface) *Worker {
	return &Worker{
		config: config,
		logger: logger,
	}
}

func (w *Worker) Run() error {
	w.logger.Info("Worker pipeline started")
	// TODO: Add actual work processing here
	return nil
}

func RunWorker() {
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

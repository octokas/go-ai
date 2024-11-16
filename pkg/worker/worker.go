package worker

import (
	"github.com/octokas/go-ai/pkg/config"
	"github.com/octokas/go-ai/pkg/logger"
)

type Worker struct {
	logger logger.LoggerInterface
}

func NewWorker(logger logger.LoggerInterface) *Worker {
	return &Worker{
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

	// Check if configuration loads successfully
	if _, err := config.Load(); err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	// Initialize and run worker
	worker := NewWorker(log)
	if err := worker.Run(); err != nil {
		log.Fatal("Worker pipeline failed:", err)
	}
}

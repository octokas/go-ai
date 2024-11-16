package worker

import (
	"testing"

	"github.com/octokas/go-ai/pkg/logger"
	"github.com/stretchr/testify/assert"
)

func TestWorker(t *testing.T) {
	t.Run("worker runs successfully", func(t *testing.T) {
		// Setup
		log := logger.New()
		w := NewWorker(log)

		// Act
		err := w.Run()

		// Assert
		assert.NoError(t, err)
	})

	t.Run("worker processes multiple items", func(t *testing.T) {
		// Setup
		log := logger.New()
		w := NewWorker(log)

		// Add test data
		testData := []string{"item1", "item2", "item3"}

		// Act & Assert
		for _, item := range testData {
			err := w.Run() // In a real scenario, you might pass the item to process
			assert.NoError(t, err, "Failed to process item: %s", item)
		}
	})
}

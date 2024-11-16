package worker

import (
	"testing"

	"github.com/octokas/go-ai/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewWorker(t *testing.T) {
	mockConfig := new(mocks.MockConfig)
	mockLogger := new(mocks.MockLogger)

	worker := NewWorker(mockConfig, mockLogger)

	assert.NotNil(t, worker)
	assert.Same(t, mockConfig, worker.config)
	assert.Same(t, mockLogger, worker.logger)
}

func TestWorkerRun(t *testing.T) {
	mockConfig := new(mocks.MockConfig)
	mockLogger := new(mocks.MockLogger)

	// Set expectations
	mockLogger.On("Info", mock.Anything).Return()

	worker := NewWorker(mockConfig, mockLogger)

	err := worker.Run()
	assert.NoError(t, err)

	// Verify expectations
	mockLogger.AssertExpectations(t)
}

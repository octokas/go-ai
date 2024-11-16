package api

import (
	"testing"

	"github.com/octokas/go-ai/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestConnectAPI(t *testing.T) {
	// Create mocks
	mockLogger := new(mocks.MockLogger)
	mockConfig := new(mocks.MockConfig)
	mockServer := new(mocks.MockServer)

	// Set expectations
	mockLogger.On("Info", mock.Anything).Return()
	mockConfig.On("GetServerConfig").Return(map[string]interface{}{
		"port": 8080,
		"host": "localhost",
	})
	mockServer.On("Start").Return(nil)

	// Override dependencies
	originalLogger := logger.New
	originalConfig := config.Load
	originalServer := server.New
	defer func() {
		logger.New = originalLogger
		config.Load = originalConfig
		server.New = originalServer
	}()

	logger.New = func() Logger {
		return mockLogger
	}
	config.Load = func() (*config.Config, error) {
		return mockConfig, nil
	}
	server.New = func(cfg *config.Config) Server {
		return mockServer
	}

	// Run the API
	ConnectAPI()

	// Verify expectations
	mockLogger.AssertExpectations(t)
	mockConfig.AssertExpectations(t)
	mockServer.AssertExpectations(t)
}

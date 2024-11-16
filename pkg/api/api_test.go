package api

import (
	"testing"

	"github.com/octokas/go-ai/pkg/config"
	"github.com/octokas/go-ai/pkg/logger"
	"github.com/octokas/go-ai/pkg/mocks"
	"github.com/octokas/go-ai/pkg/server"
	"github.com/stretchr/testify/mock"
)

// Add these type definitions at the top of your test file
// type loggerFactory func() Logger
// type configFactory func() (*config.Config, error)
// type serverFactory func(*config.Config) Server

type Server interface {
	Start() error
}

type Logger interface {
	Info(args ...interface{})
}

// Add this type to match the expected config structure
type Config struct {
	mock.Mock
}

// Replace the direct package references with variables
var (
	loggerFn = logger.New
	configFn = config.Load
	serverFn = server.New
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

	// Store original values
	originalLogger := loggerFn
	originalConfig := configFn
	originalServer := serverFn
	defer func() {
		loggerFn = originalLogger
		configFn = originalConfig
		serverFn = originalServer
	}()

	// loggerFn = func() Logger {
	// 	return mockLogger
	// }
	// configFn = func() (*config.Config, error) {
	// 	return mockConfig, nil
	// }
	// serverFn = func(cfg *config.Config) Server {
	// 	return mockServer
	// }

	// Run the API
	RunAPI()

	// Verify expectations
	mockLogger.AssertExpectations(t)
	mockConfig.AssertExpectations(t)
	mockServer.AssertExpectations(t)
}

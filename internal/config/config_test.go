package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	// Create a temporary config file
	configJSON := `{
		"server": {
			"port": 8080,
			"host": "localhost",
			"timeout": 30
		},
		"database": {
			"host": "localhost",
			"port": 5432,
			"user": "test",
			"password": "test",
			"dbname": "testdb"
		},
		"ai": {
			"model_path": "/models/test",
			"threshold": 0.5,
			"max_batch_size": 32
		}
	}`

	err := os.WriteFile("config.json", []byte(configJSON), 0644)
	assert.NoError(t, err)
	defer os.Remove("config.json")

	// Test loading config
	cfg, err := Load()
	assert.NoError(t, err)
	assert.NotNil(t, cfg)

	// Test values
	assert.Equal(t, 8080, cfg.Server.Port)
	assert.Equal(t, "localhost", cfg.Server.Host)
	assert.Equal(t, 30, cfg.Server.Timeout)

	// Test environment override
	os.Setenv("SERVER_HOST", "127.0.0.1")
	cfg, err = Load()
	assert.NoError(t, err)
	assert.Equal(t, "127.0.0.1", cfg.Server.Host)
}

func TestConfigSingleton(t *testing.T) {
	cfg1, _ := Load()
	cfg2, _ := Load()
	assert.Same(t, cfg1, cfg2, "Config should be a singleton")
} 
package config

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	// Reset the singleton before each test
	Reset()

	t.Run("loads default config from file", func(t *testing.T) {
		configJSON := `{
			"server": {
				"port": 8080,
				"host": "localhost",
				"timeout": 30
			}
		}`

		err := os.WriteFile("config.json", []byte(configJSON), 0644)
		assert.NoError(t, err)
		defer os.Remove("config.json")

		cfg, err := Load()
		assert.NoError(t, err)
		assert.Equal(t, "localhost", cfg.Server.Host)
		assert.Equal(t, 8080, cfg.Server.Port)
	})

	t.Run("environment variables override file config", func(t *testing.T) {
		Reset()

		os.Setenv("SERVER_HOST", "127.0.0.1")
		defer os.Unsetenv("SERVER_HOST")

		cfg, err := Load()
		assert.NoError(t, err)
		assert.Equal(t, "127.0.0.1", cfg.Server.Host)
	})

	t.Run("can load from reader", func(t *testing.T) {
		configJSON := `{
			"server": {
				"port": 8080,
				"host": "localhost",
				"timeout": 30
			}
		}`

		cfg, err := LoadFromReader(strings.NewReader(configJSON))
		assert.NoError(t, err)
		assert.Equal(t, "localhost", cfg.Server.Host)
		assert.Equal(t, 8080, cfg.Server.Port)
	})
}

func TestConfigSingleton(t *testing.T) {
	Reset()

	cfg1, _ := Load()
	cfg2, _ := Load()
	assert.Same(t, cfg1, cfg2, "Config should be a singleton")
}

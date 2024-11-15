package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	tests := []struct {
		name     string
		envVars  map[string]string
		expected Config
		wantErr  bool
	}{
		{
			name: "valid configuration",
			envVars: map[string]string{
				"PORT":      "8080",
				"ENV":       "development",
				"LOG_LEVEL": "debug",
			},
			expected: Config{
				Port:     8080,
				Env:      "development",
				LogLevel: "debug",
			},
			wantErr: false,
		},
		{
			name: "invalid port",
			envVars: map[string]string{
				"PORT": "invalid",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			for k, v := range tt.envVars {
				os.Setenv(k, v)
			}
			defer func() {
				for k := range tt.envVars {
					os.Unsetenv(k)
				}
			}()

			// Test
			cfg, err := LoadConfig()
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if cfg.Port != tt.expected.Port {
					t.Errorf("Port = %v, want %v", cfg.Port, tt.expected.Port)
				}
				// Add more field comparisons as needed
			}
		})
	}
}

package logger

import (
	"bytes"
	"log"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {
	// Capture output
	var buf bytes.Buffer
	logger := New()
	logger.logger = log.New(&buf, "", log.LstdFlags)

	tests := []struct {
		name      string
		logFunc   func(v ...interface{})
		level     Level
		message   string
		expected  string
		shouldLog bool
	}{
		{
			name:      "Info message with INFO level",
			logFunc:   logger.Info,
			level:     INFO,
			message:   "test info message",
			expected:  "[INFO]",
			shouldLog: true,
		},
		{
			name:      "Debug message with INFO level",
			logFunc:   logger.Debug,
			level:     INFO,
			message:   "test debug message",
			expected:  "[DEBUG]",
			shouldLog: false,
		},
		{
			name:      "Error message with INFO level",
			logFunc:   logger.Error,
			level:     INFO,
			message:   "test error message",
			expected:  "[ERROR]",
			shouldLog: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf.Reset()
			logger.SetLevel(tt.level)
			tt.logFunc(tt.message)

			if tt.shouldLog {
				assert.True(t, strings.Contains(buf.String(), tt.expected))
				assert.True(t, strings.Contains(buf.String(), tt.message))
			} else {
				assert.Empty(t, buf.String())
			}
		})
	}
}

func TestLoggerSingleton(t *testing.T) {
	logger1 := New()
	logger2 := New()
	assert.Same(t, logger1, logger2, "Logger should be a singleton")
}

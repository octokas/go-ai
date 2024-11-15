package logger

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLoggerMiddleware(t *testing.T) {
	// Capture log output
	var buf bytes.Buffer
	log.SetOutput(&buf)

	tests := []struct {
		name           string
		method         string
		path           string
		expectedStatus int
		expectedLog    string
	}{
		{
			name:           "successful request",
			method:         "GET",
			path:           "/test",
			expectedStatus: http.StatusOK,
			expectedLog:    "GET /test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create test handler
			nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.expectedStatus)
			})

			// Create request and response recorder
			req := httptest.NewRequest(tt.method, tt.path, nil)
			w := httptest.NewRecorder()

			// Execute middleware
			LoggerMiddleware(nextHandler).ServeHTTP(w, req)

			// Check status code
			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			// Check log output
			logOutput := buf.String()
			if !strings.Contains(logOutput, tt.expectedLog) {
				t.Errorf("Expected log to contain %q, got %q", tt.expectedLog, logOutput)
			}

			// Clear buffer for next test
			buf.Reset()
		})
	}
}

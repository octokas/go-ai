package server

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/octokas/go-ai/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	// Create test config
	cfg := &config.Config{}
	cfg.Server.Port = 8080
	cfg.Server.Host = "localhost"
	cfg.Server.Timeout = 30

	// Create server
	srv := New(cfg)
	assert.NotNil(t, srv)

	// Test health check endpoint
	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	healthCheckHandler(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "OK", w.Body.String())

	// Test server shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := srv.Shutdown(ctx)
	assert.NoError(t, err)
}

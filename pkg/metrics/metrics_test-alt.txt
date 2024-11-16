package metrics

import (
	"testing"
	"time"
)

func TestNewMetricsCollector(t *testing.T) {
	collector := NewMetricsCollector()
	if collector == nil {
		t.Error("Expected non-nil MetricsCollector")
	}
}

func TestMetricsCollector_RecordRequest(t *testing.T) {
	tests := []struct {
		name       string
		path       string
		method     string
		statusCode int
		duration   time.Duration
	}{
		{
			name:       "successful GET request",
			path:       "/api/v1/users",
			method:     "GET",
			statusCode: 200,
			duration:   100 * time.Millisecond,
		},
		{
			name:       "failed POST request",
			path:       "/api/v1/users",
			method:     "POST",
			statusCode: 400,
			duration:   50 * time.Millisecond,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			collector := NewMetricsCollector()
			collector.RecordRequest(tt.path, tt.method, tt.statusCode, tt.duration)

			// Verify metrics were recorded correctly
			metrics := collector.GetMetrics()
			if metrics.TotalRequests <= 0 {
				t.Error("Expected TotalRequests to be incremented")
			}
		})
	}
}

package metrics

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMetrics(t *testing.T) {
	metrics := New()

	// Test counter
	metrics.IncrementCounter("test_counter")
	assert.Equal(t, int64(1), metrics.GetCounter("test_counter"))
	metrics.IncrementCounter("test_counter")
	assert.Equal(t, int64(2), metrics.GetCounter("test_counter"))

	// Test timer
	duration := time.Second
	metrics.RecordTime("test_timer", duration)
	assert.Equal(t, duration, metrics.GetTimer("test_timer"))
}

func TestMetricsSingleton(t *testing.T) {
	metrics1 := New()
	metrics2 := New()
	assert.Same(t, metrics1, metrics2, "Metrics should be a singleton")
}

package pipeline

import (
	"context"
	"testing"

	"github.com/octokas/go-ai/pkg/ai"
	"github.com/octokas/go-ai/pkg/config"
	"github.com/stretchr/testify/assert"
)

// Mock step for testing
type MockStep struct {
	ExecuteFunc func(ctx context.Context, data interface{}) (interface{}, error)
}

func (m *MockStep) Execute(ctx context.Context, data interface{}) (interface{}, error) {
	return m.ExecuteFunc(ctx, data)
}

func TestPipeline(t *testing.T) {
	cfg := &config.Config{}
	model := ai.NewModel(cfg)
	pipeline := New(model)
	assert.NotNil(t, pipeline)

	// Add mock step
	mockStep := &MockStep{
		ExecuteFunc: func(ctx context.Context, data interface{}) (interface{}, error) {
			// Simply return the input data
			return data, nil
		},
	}
	pipeline.AddStep(mockStep)

	// Test pipeline execution
	ctx := context.Background()
	inputData := "test data"
	result, err := pipeline.Run(ctx, inputData)

	assert.NoError(t, err)
	assert.Equal(t, inputData, result)
}

package ai

import (
	"testing"

	"github.com/octokas/go-ai/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestModel(t *testing.T) {
	cfg := &config.Config{}
	cfg.AI.ModelPath = "/path/to/model"
	cfg.AI.Threshold = 0.5
	cfg.AI.MaxBatchSize = 32

	model := NewModel(cfg)
	assert.NotNil(t, model)

	// Test prediction
	input := []float64{1.0, 2.0, 3.0}
	output, err := model.Predict(input)
	assert.Nil(t, output) // Since we haven't implemented real prediction yet
	assert.Nil(t, err)

	// Test training
	data := [][]float64{{1.0, 2.0}, {3.0, 4.0}}
	labels := []float64{0, 1}
	err = model.Train(data, labels)
	assert.Nil(t, err)
}

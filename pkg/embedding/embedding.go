package embedding

import (
	"fmt"
)

// EmbeddingService defines the interface for embedding providers
type EmbeddingService interface {
	GetEmbeddings(text string) ([]float32, error)
	Close() error
}

// NewEmbeddingService creates a new embedding service based on configuration
func NewEmbeddingService(config EmbeddingConfig) (EmbeddingService, error) {
	if err := config.Validate(); err != nil {
		return nil, err
	}

	switch config.Provider {
	case "openai":
		return NewOpenAIEmbedding(config.APIKey, config.ModelName), nil
	default:
		return nil, fmt.Errorf("unsupported embedding provider: %s", config.Provider)
	}
} 
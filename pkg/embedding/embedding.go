package embedding

import (
	"fmt"
	"os"
)

// EmbeddingService defines the interface for embedding providers
type EmbeddingService interface {
	// Provider  string
	// ModelName string
	GetEmbeddings(text string) ([]float32, error)
	Close() error
}

// NewEmbeddingService creates a new embedding service based on configuration
func NewEmbeddingService() (EmbeddingService, error) {
	provider := os.Getenv("EMBEDDING_PROVIDER")
	model := os.Getenv("EMBEDDING_MODEL")
	switch provider {
	case "local":
		if model == "" {
			return nil, fmt.Errorf("MODEL_PATH is required when EMBEDDING_PROVIDER is local")
		}
		return NewLocalEmbedding(model), nil
	default:
		return nil, fmt.Errorf("unsupported embedding provider: %s", provider)
	}
}

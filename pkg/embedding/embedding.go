package embedding

import (
	"fmt"
	"os"
	"strings"
)

// type EmbeddingConfig struct {
// 	Provider  string // "local" or other providers in the future
// 	ModelName string // path to local model
// }

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
	case "local":
		return NewLocalEmbedding(config.ModelName), nil
	default:
		return nil, fmt.Errorf("unsupported embedding provider: %s", config.Provider)
	}
}

func Validate(c EmbeddingConfig) error {
	provider := os.Getenv("EMBEDDING_PROVIDER")
	model := os.Getenv("EMBEDDING_MODEL")
	switch provider {
	case "local":
		provider = strings.ToLower(provider)
		if c.ModelName == "" {
			return fmt.Errorf("MODEL_PATH is required when EMBEDDING_PROVIDER is local")
		}
	default:
		return fmt.Errorf("unsupported embedding provider: %s", c.Provider)
	}

	if model == "" {
		c.ModelName = "/pkg/ai/models/llama3-8b-instruct.Q4_0.gguf" // set default model path
		return nil                                                  // return nil error since we've set a default
	}
	c.ModelName = model
	return nil
}

package embedding

import (
	"fmt"
	"os"
	"strings"
)

type EmbeddingConfig struct {
	Provider  string // "openai" or other providers in the future
	APIKey    string
	ModelName string // optional, defaults per provider
}

func LoadEmbeddingConfig() EmbeddingConfig {
	return EmbeddingConfig{
		Provider:  getEmbeddingProvider(),
		APIKey:    os.Getenv("OPENAI_API_KEY"),
		ModelName: getModelName(),
	}
}

func (c EmbeddingConfig) Validate() error {
	switch c.Provider {
	case "openai":
		if c.APIKey == "" {
			return fmt.Errorf("OPENAI_API_KEY is required when EMBEDDING_PROVIDER is openai")
		}
	default:
		return fmt.Errorf("unsupported embedding provider: %s", c.Provider)
	}
	return nil
}

func getEmbeddingProvider() string {
	provider := os.Getenv("EMBEDDING_PROVIDER")
	if provider == "" {
		return "openai" // default provider
	}
	return strings.ToLower(provider)
}

func getModelName() string {
	model := os.Getenv("EMBEDDING_MODEL")
	if model == "" {
		return "text-embedding-ada-002" // default model
	}
	return model
}

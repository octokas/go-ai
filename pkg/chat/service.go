package chat

import (
	"fmt"
	"strings"

	"github.com/octokas/go-ai/pkg/vectorstore"
)

type Service struct {
	vectorStore vectorstore.Store
	embedder    EmbeddingService
	llm         LLMService
	config      ServiceConfig
}

type ServiceConfig struct {
	MaxContextDocs      int
	MaxTokensPerDoc     int
	SimilarityThreshold float32
}

type SearchResult struct {
	vectorstore.Document
	Score float32
}

type EmbeddingService interface {
	GetEmbeddings(text string) ([]float32, error)
}

type LLMService interface {
	Complete(prompt string) (string, error)
}

func (s *Service) ProcessMessage(message string) (string, error) {
	// 1. Search for relevant documents
	results, err := s.vectorStore.Search(message, s.config.MaxContextDocs)
	if err != nil {
		return "", fmt.Errorf("search failed: %w", err)
	}

	// 2. Filter and prepare context
	filteredDocs := s.filterResults(results)
	prompt := s.constructPrompt(message, filteredDocs)

	// 3. Get LLM response
	response, err := s.llm.Complete(prompt)
	if err != nil {
		return "", fmt.Errorf("llm completion failed: %w", err)
	}

	return response, nil
}

func (s *Service) filterResults(results []vectorstore.SearchResult) []vectorstore.Document {
	filtered := make([]vectorstore.Document, 0)
	totalTokens := 0

	for _, result := range results {
		if result.Score < s.config.SimilarityThreshold {
			continue
		}

		// Estimate tokens (you'll need a proper token counting function)
		docTokens := estimateTokens(result.Content)
		if totalTokens+docTokens > s.config.MaxTokensPerDoc {
			break
		}

		filtered = append(filtered, result.Document)
		totalTokens += docTokens
	}

	return filtered
}

func (s *Service) constructPrompt(message string, docs []vectorstore.Document) string {
	context := ""
	for _, doc := range docs {
		context += doc.Content + "\n"
	}
	return fmt.Sprintf("Context:\n%s\nQuestion: %s", context, message)
}

// estimateTokens provides a rough estimation of tokens in a text
// A simple approximation is counting words (splitting by whitespace)
func estimateTokens(text string) int {
	return len(strings.Fields(text))
}

func NewService(store vectorstore.Store, embedder EmbeddingService, llm LLMService, config ServiceConfig) *Service {
	return &Service{
		vectorStore: store,
		embedder:    embedder,
		llm:         llm,
		config:      config,
	}
}

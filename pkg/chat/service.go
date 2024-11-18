package chat

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/octokas/go-ai/pkg/embedding"
	"github.com/octokas/go-ai/pkg/llm"
	"github.com/octokas/go-ai/pkg/memory"
)

type Service struct {
	store      *memory.MemoryStore
	embedder   embedding.EmbeddingService
	llm        llm.LLMService
	config     ChatConfig
	memory     MessageHistory
	middleware []ChatMiddleware
}

type ServiceOptions struct {
	Store      *memory.MemoryStore
	Embedder   embedding.EmbeddingService
	LLM        llm.LLMService
	Config     ChatConfig
	Middleware []ChatMiddleware // For logging, rate limiting, etc.
}

// ChatMiddleware for extensibility
type ChatMiddleware func(ChatHandler) ChatHandler
type ChatHandler func(context.Context, *ChatRequest) (*ChatResponse, error)

type ChatRequest struct {
	Message   string
	SessionID string
	UserID    string
	Metadata  map[string]interface{}
}

type Source struct {
	Content    string
	Reference  string
	Similarity float64
}

type ChatResponse struct {
	Message    string
	Sources    []Source
	TokensUsed int
	Metadata   map[string]interface{}
}

// MessageHistory represents the chat history management interface
type MessageHistory interface {
	// Add basic methods as needed
	AddMessage(message string) error
	GetMessages() []string
}

// NewMessageHistory creates a new message history with the specified size
func NewMessageHistory(size int) MessageHistory {
	// TODO: Implement actual message history storage
	return &memoryHistory{
		messages: make([]string, 0, size),
		maxSize:  size,
	}
}

type memoryHistory struct {
	messages []string
	maxSize  int
}

func (h *memoryHistory) AddMessage(message string) error {
	if len(h.messages) >= h.maxSize {
		h.messages = h.messages[1:]
	}
	h.messages = append(h.messages, message)
	return nil
}

func (h *memoryHistory) GetMessages() []string {
	return h.messages
}

func NewService(opts ServiceOptions) *Service {
	svc := &Service{
		store:      opts.Store,
		embedder:   opts.Embedder,
		llm:        opts.LLM,
		config:     opts.Config,
		memory:     NewMessageHistory(opts.Config.MemorySize),
		middleware: opts.Middleware,
	}

	return svc
}

func (s *Service) baseChatHandler(ctx context.Context, req *ChatRequest) (*ChatResponse, error) {
	// Basic implementation - you'll want to expand this
	return &ChatResponse{
		Message: "Default response", // Replace with actual LLM call
	}, nil
}

func (s *Service) Chat(ctx context.Context, req *ChatRequest) (*ChatResponse, error) {
	// Apply middleware chain
	handler := s.baseChatHandler
	for i := len(s.middleware) - 1; i >= 0; i-- {
		handler = s.middleware[i](handler)
	}

	return handler(ctx, req)
}

func ChatServer(service *Service) error {
	r := gin.Default()

	r.POST("/chat", func(c *gin.Context) {
		var req ChatRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		response, err := service.Chat(c.Request.Context(), &req)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"response": response})
	})

	return r.Run(":4040")
}

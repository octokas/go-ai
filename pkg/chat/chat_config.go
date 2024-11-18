package chat

type ChatConfig struct {
	MaxContextDocs      int
	MaxTokensPerDoc     int
	SimilarityThreshold float64
	Temperature         float64
	MaxResponseTokens   int
	Model               string
	MemorySize          int
	ContextWindow       int
	RequestsPerMinute   int
	ConcurrentChats     int
}

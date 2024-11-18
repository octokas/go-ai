package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/octokas/go-ai/pkg/chat"
	"github.com/octokas/go-ai/pkg/embedding"
	"github.com/octokas/go-ai/pkg/router"
	vectorstore "github.com/octokas/go-ai/pkg/vectorstore"
)

func HelloDutonian() string {
	return "Hello, Dutonian! 👋🏼"
}

func main() {
	// Load and validate vector store configuration
	vectorConfig := vectorstore.LoadVectorStoreConfig()
	if err := vectorConfig.Validate(); err != nil {
		log.Fatalf("Vector store configuration error: %v", err)
	}

	// Load and validate embedding configuration
	embeddingConfig := embedding.LoadEmbeddingConfig()
	if err := embeddingConfig.Validate(); err != nil {
		log.Fatalf("Embedding configuration error: %v", err)
	}

	// Initialize vector store
	store, err := vectorstore.NewStore(vectorConfig)
	if err != nil {
		log.Fatalf("Failed to create vector store: %v", err)
	}
	defer store.Close()

	// Initialize embedding service
	embedder, err := embedding.NewEmbeddingService(embeddingConfig)
	if err != nil {
		log.Fatalf("Failed to create embedding service: %v", err)
	}
	defer embedder.Close()

	// Initialize LLM service
	llm := chat.NewLLMService()

	service := chat.NewService(
		store,
		embedder,
		llm,
		chat.ServiceConfig{
			MaxContextDocs:      5,
			MaxTokensPerDoc:     1000,
			SimilarityThreshold: 0.7,
		},
	)

	// Define flags for different servers
	apiV1 := flag.Bool("apiv1", false, "Start only APIv1 server")
	apiV2 := flag.Bool("apiv2", false, "Start only APIv2 server")
	home := flag.Bool("home", false, "Start only home server")
	chat := flag.Bool("chat", false, "Start only chat server")
	flag.Parse()

	// Channel to collect errors from all components
	errorChan := make(chan error, 10)
	//done := make(chan bool)

	// Error monitoring goroutine
	go func() {
		for err := range errorChan {
			log.Printf("ERROR: %+v\n", err)
		}
	}()

	fmt.Println(HelloDutonian())

	// Setup routes based on flags
	if !*apiV2 && !*home && !*apiV1 && !*chat {
		log.Println("Starting all servers...")
		go router.HomeServer()        // 8080
		go router.V1Server()          // 2020
		go router.V2Server()          // 3030
		go router.ChatServer(service) // Pass service to ChatServer 4040
	} else if *apiV2 {
		go router.V2Server()
	} else if *home {
		go router.HomeServer()
	} else if *apiV1 {
		go router.V1Server()
	} else if *chat {
		go router.ChatServer(service)
	} else {
		log.Println("No server type specified. Use -all flag to start all servers")
		return
	}

	// Block forever to keep the program running
	select {}

}

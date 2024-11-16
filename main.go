package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	api "github.com/octokas/go-ai/pkg/api"
	worker "github.com/octokas/go-ai/pkg/worker"
	changelog "github.com/octokas/go-ai/scripts/changelog"
	fix_yaml_comments "github.com/octokas/go-ai/scripts/fix_yaml_comments"
	run_tests "github.com/octokas/go-ai/tests/run_tests"
)

func main() {
	// Create WaitGroup to track all running components
	var wg sync.WaitGroup

	// Setup signal handling and shutdown channel
	sigChan := make(chan os.Signal, 1)
	shutdown := make(chan struct{})
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Start API server
	wg.Add(1)
	go func() {
		defer wg.Done()
		done := make(chan struct{})
		go func() {
			api.RunAPI()
			close(done)
		}()

		select {
		case <-done:
			log.Println("API server stopped")
		case <-shutdown:
			log.Println("API server shutdown requested")
		}
	}()

	// Start worker
	wg.Add(1)
	go func() {
		defer wg.Done()
		done := make(chan struct{})
		errChan := make(chan error, 1)

		go func() {
			worker.RunWorker()
			close(done)
		}()

		select {
		case err := <-errChan:
			log.Printf("Worker stopped with error: %v", err)
		case <-done:
			log.Println("Worker completed normally")
		case <-shutdown:
			log.Println("Worker shutdown requested")
		}
	}()

	// Run other components
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-shutdown:
			return
		default:
			changelog.GenerateChangelog()
			fix_yaml_comments.FixYAMLComments(".")
			run_tests.RunTests()
			fmt.Println(HelloDutonian())
		}
	}()

	// Wait for shutdown signal and trigger shutdown
	sig := <-sigChan
	log.Printf("Received signal: %v", sig)
	close(shutdown)

	// Wait for all components with timeout
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		log.Println("All components shut down successfully")
	case <-time.After(10 * time.Second):
		log.Println("Shutdown timed out")
	}
}

func HelloDutonian() string {
	return "Hello, Dutonian! 👋🏼"
}

// Template for adding new components with graceful shutdown
/*
   // Start [COMPONENT_NAME]
   wg.Add(1)
   go func() {
       defer wg.Done()
       done := make(chan struct{})
       go func() {
           // Replace this line with your component's run function
           // example.RunComponent()
           close(done)
       }()

       select {
       case <-done:
           log.Println("[COMPONENT_NAME] stopped")
       case <-shutdown:
           log.Println("[COMPONENT_NAME] shutdown requested")
       }
   }()
*/

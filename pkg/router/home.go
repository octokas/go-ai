package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kyokomi/emoji/v2"
	"github.com/octokas/go-ai/pkg/config"
	"github.com/octokas/go-ai/pkg/logger"
	"github.com/octokas/go-ai/pkg/server"
	"github.com/octokas/go-ai/pkg/utils"
)

func setupHomeRoutes() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contactHandler)
}

// USING FPRINTF
// func homeHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Welcome new Dutonian! 📟🎉")
// }

// JUST USING EMOJI
// func homeHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html; charset=utf-8") // Important for emoji support
// 	content := emoji.Sprint(`
//         <h1>Welcome new Dutonian! :pager: :tada:</h1>
//         <p>Need help getting started? <a href="/chat/prompt">Chat with our Dutonian Assistant</a> :robot:</p>
//     `)
// 	fmt.Fprint(w, content)
// }

// USING UTILS.WRITEHTML
func homeHandler(w http.ResponseWriter, r *http.Request) {
	utils.WriteHTML(w, `
        <h1>Welcome new Dutonian! :pager: :tada:</h1>
        <p>Need help getting started? <a href="http://localhost:4040/chat/prompt">Chat with our Dutonian Assistant</a> :robot:</p>
    `)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	content := emoji.Sprint("Contact Us: hello@dutonian.com :email:")
	fmt.Fprint(w, content)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About Us: We're building something awesome! 🚀")
}

// func contactHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Contact Us: hello@dutonian.com 📧")
// }

func HomeServer() { // Initial logging with standard log package
	log.Println("Initializing application...")

	// Initialize logger
	logger := logger.New()
	logger.Info("Starting Home server...")

	// Setup home routes
	setupHomeRoutes()

	// Start server
	log.Printf("[INFO] Starting server on :8080")
	http.ListenAndServe(":8080", nil)

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		logger.Fatal("Failed to load configuration:", err)
	}

	// Initialize server
	srv := server.New(cfg)

	// Start server
	if err := srv.Start(); err != nil && err != http.ErrServerClosed {
		logger.Fatal("Server failed:", err)
	}
}

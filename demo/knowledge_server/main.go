package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/octokas/go-ai/knowledge/api"
	"github.com/octokas/go-ai/knowledge/library"
	"github.com/octokas/go-ai/knowledge/storage"
)

func main() {
	// Create storage directory in the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	storagePath := filepath.Join(cwd, "knowledge_data")

	// Initialize storage
	store, err := storage.NewFileStorage(storagePath)
	if err != nil {
		log.Fatal("Failed to initialize storage:", err)
	}

	// Initialize library
	lib := library.NewLibrary(store)
	if err := lib.LoadAllDocuments(); err != nil {
		log.Fatal("Failed to load documents:", err)
	}

	// Initialize handlers
	uploadHandler := api.NewUploadHandler(lib)

	// Set up routes
	http.HandleFunc("/upload", uploadHandler.HandleUpload)
	//http.HandleFunc("/", handleHome)

	// Add to your routes in main.go
	http.HandleFunc("/chat", api.NewChatHandler(lib))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		http.ServeFile(w, r, "templates/chat.html")
	})

	// Start server
	port := ":1313"
	fmt.Printf("Starting knowledge server on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	html := `
<!DOCTYPE html>
<html>
<head>
    <title>Knowledge Library Upload</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            max-width: 800px;
            margin: 0 auto;
            padding: 20px;
        }
        .upload-form {
            border: 1px solid #ccc;
            padding: 20px;
            border-radius: 5px;
        }
        .button {
            background-color: #4CAF50;
            border: none;
            color: white;
            padding: 15px 32px;
            text-align: center;
            text-decoration: none;
            display: inline-block;
            font-size: 16px;
            margin: 4px 2px;
            cursor: pointer;
            border-radius: 4px;
        }
    </style>
</head>
<body>
    <h1>Knowledge Library Upload</h1>
    <div class="upload-form">
        <form action="/upload" method="post" enctype="multipart/form-data">
            <p>Select a text file to upload:</p>
            <input type="file" name="document" accept=".txt,.md,.json">
            <br><br>
            <input type="submit" value="Upload" class="button">
        </form>
    </div>
</body>
</html>
`
	fmt.Fprint(w, html)
}

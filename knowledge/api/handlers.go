package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/octokas/go-ai/knowledge/library"
)

type UploadHandler struct {
	lib *library.Library
}

// type ChatHandler struct {
// 	lib *library.Library
// }

// PLAIN UPLOADER
// func NewChatHandler(lib *library.Library) *ChatHandler {
// 	return &ChatHandler{lib: lib}
// }

// CHAT HANDLER with fancy prompt
// func NewChatHandler(lib *library.Library) *ChatHandler {
// 	return &ChatHandler{lib: lib}
// }

func NewChatHandler(lib *library.Library) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var request struct {
			Message string `json:"message"`
		}

		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		// Search the knowledge base
		results := lib.Search(request.Message)

		// Prepare response
		response := "Based on our knowledge base: "
		if len(results) > 0 {
			response += strings.Join(results, "\n\n")
		} else {
			response = "I couldn't find any relevant information in the knowledge base."
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"response": response,
		})
	}
}

func NewUploadHandler(lib *library.Library) *UploadHandler {
	return &UploadHandler{lib: lib}
}

func (h *UploadHandler) HandleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	file, header, err := r.FormFile("document")
	if err != nil {
		http.Error(w, "Error reading file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Read file content
	content := make([]byte, header.Size)
	if _, err := file.Read(content); err != nil {
		http.Error(w, "Error reading file content", http.StatusInternalServerError)
		return
	}

	// Add to library
	err = h.lib.AddDocument(header.Filename, string(content))
	if err != nil {
		http.Error(w, "Error adding document to library", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

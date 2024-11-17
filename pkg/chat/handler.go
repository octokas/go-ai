package chat

import (
	"encoding/json"
	"net/http"
)

var service *Service

func init() {
	service = NewService()
}

type Request struct {
	Message string `json:"message"`
}

type Response struct {
	Response string `json:"response"`
}

func HandleAPI(w http.ResponseWriter, r *http.Request) {
	var req Request
	var resp string
	var err error

	switch {
	case r.Method != http.MethodPost:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return

	case json.NewDecoder(r.Body).Decode(&req) != nil:
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return

	case func() bool {
		resp, err = service.ProcessMessage(req.Message)
		return err != nil
	}():
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return

	default:
		response := Response{
			Response: "Hello! I'm Claude, an AI assistant. I'm here to help you. " + resp,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

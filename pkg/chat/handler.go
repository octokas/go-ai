package chat

import (
	"encoding/json"
	"net/http"

	"github.com/kyokomi/emoji/v2"
)

type Handler struct {
	service  *Service
	greeting string
}

type Request struct {
	Message string `json:"message"`
}

type Response struct {
	Response string `json:"response"`
	Error    string `json:"error,omitempty"`
}

func (s *Service) ProcessMessage(message string) (string, error) {
	// Add your message processing logic here
	return "Response to: " + message, nil
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service:  service,
		greeting: emoji.Sprint("Hello! I'm PagerMate :pager:, your Dutonian AI assistant. I'm here to help you get the most out of your Dutonian onboarding experience! :rocket: "),
	}
}

func (h *Handler) HandleAPI(w http.ResponseWriter, r *http.Request) {
	// Only allow POST requests
	if r.Method != http.MethodPost {
		h.writeError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse request
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.writeError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate request
	if req.Message == "" {
		h.writeError(w, "Message is required", http.StatusBadRequest)
		return
	}

	// Process message
	resp, err := h.service.ProcessMessage(req.Message)
	if err != nil {
		h.writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send response
	response := Response{
		Response: h.greeting + resp,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) writeError(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(Response{
		Error: message,
	})
}

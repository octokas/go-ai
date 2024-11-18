package embedding

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type OpenAIEmbedding struct {
	apiKey string
	model  string
	client *http.Client
}

func NewOpenAIEmbedding(apiKey string, model string) *OpenAIEmbedding {
	if model == "" {
		model = "text-embedding-ada-002"
	}
	
	return &OpenAIEmbedding{
		apiKey: apiKey,
		model:  model,
		client: &http.Client{},
	}
}

func (e *OpenAIEmbedding) GetEmbeddings(text string) ([]float32, error) {
	url := "https://api.openai.com/v1/embeddings"
	
	requestBody := map[string]interface{}{
		"input": text,
		"model": e.model,
	}
	
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+e.apiKey)

	resp, err := e.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(body))
	}

	var result struct {
		Data []struct {
			Embedding []float32 `json:"embedding"`
		} `json:"data"`
		Error *struct {
			Message string `json:"message"`
		} `json:"error"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if result.Error != nil {
		return nil, fmt.Errorf("API error: %s", result.Error.Message)
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("no embedding returned")
	}

	return result.Data[0].Embedding, nil
}

func (e *OpenAIEmbedding) Close() error {
	// Clean up any resources if needed
	return nil
} 
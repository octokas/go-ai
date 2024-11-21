package middleware

import (
	"context"
	"net/http"
)

type CardData struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Image   string   `json:"image,omitempty"`
	Actions []Action `json:"actions,omitempty"`
}

type Action struct {
	Text  string `json:"text"`
	Class string `json:"class"`
	URL   string `json:"url,omitempty"`
}

func DynamicContent(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add dynamic content data to context
		ctx := r.Context()
		cards := []CardData{
			{
				Title:   "Featured Item",
				Content: "This is a dynamic card loaded via Go middleware",
				Actions: []Action{
					{
						Text:  "Learn More",
						Class: "btn-primary",
					},
				},
			},
		}

		ctx = context.WithValue(ctx, "cards", cards)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

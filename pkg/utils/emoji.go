package utils

import (
	"net/http"

	"github.com/kyokomi/emoji/v2"
)

func WriteHTML(w http.ResponseWriter, content string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	emoji.Fprint(w, content)
}

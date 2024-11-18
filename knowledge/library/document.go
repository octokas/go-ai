package library

import (
	"strings"
)

type Document struct {
	content   string
	words     map[string]int
	sentences []string
}

func NewDocument(content string) *Document {
	doc := &Document{
		content: content,
		words:   make(map[string]int),
	}
	doc.process()
	return doc
}

func (d *Document) process() {
	// Split content into sentences (simplified)
	d.sentences = strings.Split(d.content, ".")

	// Build word frequency map
	words := strings.Fields(strings.ToLower(d.content))
	for _, word := range words {
		d.words[word]++
	}
}

func (d *Document) MatchesQuery(queryTerms []string) bool {
	for _, term := range queryTerms {
		if _, exists := d.words[term]; !exists {
			return false
		}
	}
	return true
}

func (d *Document) GetRelevantExcerpts(queryTerms []string) []string {
	relevant := []string{}

	for _, sentence := range d.sentences {
		sentenceLower := strings.ToLower(sentence)
		for _, term := range queryTerms {
			if strings.Contains(sentenceLower, term) {
				relevant = append(relevant, strings.TrimSpace(sentence))
				break
			}
		}
	}

	return relevant
}

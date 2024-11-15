package tests

import (
	"log"

	test_reporter "github.com/octokas/go-ai/scripts/test_reporter"
)

func RunTests() {
	if err := test_reporter.SaveTestReports(); err != nil {
		log.Fatalf("Failed to save test reports: %v", err)
	}
}

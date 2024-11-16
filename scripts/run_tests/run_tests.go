package run_tests

import (
	"log"

	test_reporter "github.com/octokas/go-ai/scripts/test_reporter"
)

var testReporter test_reporter.TestReporter = &test_reporter.Reporter{}

func RunTests() error {
	if err := testReporter.SaveTestReports(); err != nil {
		log.Fatalf("Failed to save test reports: %v", err)
	}
	return testReporter.SaveTestReports()
}

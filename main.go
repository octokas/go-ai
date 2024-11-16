package main

import (
	api "github.com/octokas/go-ai/cmd/api"
	worker "github.com/octokas/go-ai/cmd/worker"
	changelog "github.com/octokas/go-ai/scripts/changelog"
	fix_yaml_comments "github.com/octokas/go-ai/scripts/fix_yaml_comments"
	run_tests "github.com/octokas/go-ai/scripts/run_tests"
	test_reporter "github.com/octokas/go-ai/scripts/test_reporter"
)

func main() {
	changelog.GenerateChangelog()
	fix_yaml_comments.FixYAMLComments(".")
	run_tests.RunTests()
	test_reporter.SaveTestReports()

	api.ConnectAPI()
	worker.SetupWorker()
}

package main

import (
	api "github.com/octokas/go-ai/pkg/api"
	worker "github.com/octokas/go-ai/pkg/worker"
	changelog "github.com/octokas/go-ai/scripts/changelog"
	fix_yaml_comments "github.com/octokas/go-ai/scripts/fix_yaml_comments"
	run_tests "github.com/octokas/go-ai/tests/run_tests"
)

func main() {
	changelog.GenerateChangelog()
	fix_yaml_comments.FixYAMLComments(".")
	run_tests.RunTests()

	api.RunAPI()
	worker.RunWorker()
}

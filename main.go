package main

import (
	"fmt"

	api "github.com/octokas/go-ai/pkg/api"
	worker "github.com/octokas/go-ai/pkg/worker"
	changelog "github.com/octokas/go-ai/scripts/changelog"
	fix_yaml_comments "github.com/octokas/go-ai/scripts/fix_yaml_comments"
	run_tests "github.com/octokas/go-ai/tests/run_tests"
)

func main() {
	api.RunAPI()
	worker.RunWorker()
	changelog.GenerateChangelog()
	fix_yaml_comments.FixYAMLComments(".")
	run_tests.RunTests()
	fmt.Println(HelloDutonian())
}

func HelloDutonian() string {
	return "Hello, Dutonian! :wave:"
}

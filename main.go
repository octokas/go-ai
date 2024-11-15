package main

import (
	changelog "github.com/octokas/go-ai/scripts/changelog"
	fix_yaml_comments "github.com/octokas/go-ai/scripts/fix_yaml_comments"
	test_reporter "github.com/octokas/go-ai/scripts/test_reporter"
	tests "github.com/octokas/go-ai/scripts/tests"
)

func main() {
	changelog.GenerateChangelog()
	fix_yaml_comments.FixYAMLComments(".")
	tests.RunTests()
	test_reporter.SaveTestReports()
}

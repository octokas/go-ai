package main

import (
	changelog "github.com/octokas/go-ai/scripts"
	fix_yaml_comments "github.com/octokas/go-ai/scripts/fix_yaml_comments"
)

func main() {
	changelog.GenerateChangelog()
	fix_yaml_comments.FixYAMLComments(".")
}

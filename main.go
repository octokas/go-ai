package main

import (
	"github.com/octokas/go-ai/scripts/changelog"
	"github.com/octokas/go-ai/scripts/yaml_comments"
)

func main() {
	changelog.GenerateChangelog()
	yaml_comments.FixYAMLComments(".")
}

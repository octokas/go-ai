package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func main() {
	// Get the latest tag
	lastTag, _ := exec.Command("git", "describe", "--tags", "--abbrev=0").Output()
	
	// Get all commits since last tag
	cmd := exec.Command("git", "log", "--pretty=format:%s", string(lastTag))
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error getting git log: %v\n", err)
		return
	}

	// Parse commits
	commits := strings.Split(string(output), "\n")
	
	// Generate changelog
	changelog := fmt.Sprintf("# Changelog\n\n## [Unreleased] - %s\n\n", time.Now().Format("2006-01-02"))
	
	features := []string{}
	fixes := []string{}
	others := []string{}

	for _, commit := range commits {
		if strings.HasPrefix(commit, "feat:") {
			features = append(features, strings.TrimPrefix(commit, "feat:"))
		} else if strings.HasPrefix(commit, "fix:") {
			fixes = append(fixes, strings.TrimPrefix(commit, "fix:"))
		} else {
			others = append(others, commit)
		}
	}

	if len(features) > 0 {
		changelog += "### Features\n\n"
		for _, f := range features {
			changelog += fmt.Sprintf("* %s\n", strings.TrimSpace(f))
		}
		changelog += "\n"
	}

	if len(fixes) > 0 {
		changelog += "### Bug Fixes\n\n"
		for _, f := range fixes {
			changelog += fmt.Sprintf("* %s\n", strings.TrimSpace(f))
		}
		changelog += "\n"
	}

	if len(others) > 0 {
		changelog += "### Other Changes\n\n"
		for _, o := range others {
			changelog += fmt.Sprintf("* %s\n", strings.TrimSpace(o))
		}
	}

	fmt.Println(changelog)
} 
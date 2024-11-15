package changelog

func generateChangelog(commitID, author string) (string, error) {
	changelog := `# Changelog

Commit: ` + commitID + `
Author: ` + author + `

## Changes
- Add your changes here
`
	return changelog, nil
}

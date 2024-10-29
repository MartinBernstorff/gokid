package version_control

import (
	"gokid/forge"
	"strings"
)

// VCS defines the interface for version control operations
type VCS interface {
	NewChange(issue forge.Issue, defaultBranch string, migrateChanges bool, branchPrefix string, branchSuffix string) error
}

// branchTitle creates a branch name from an issue title
func branchTitle(issue forge.Issue, prefix string, suffix string) string {
	title := prefix + issue.Title.Content + suffix
	return strings.ReplaceAll(title, " ", "-")
}

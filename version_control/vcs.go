package version_control

import (
	"gokid/forge"
	"strings"
)

// VCS defines the interface for version control operations
type VCS interface {
	NewChange(issue forge.Issue, defaultBranch string, migrateChanges bool, branchPrefix string, branchSuffix string) error
	SyncTrunk(defaultBranch string) error
	ShowDiffSummary(branch string) error
	IsClean() bool
}

// BranchTitle creates a branch name from an issue title
func BranchTitle(issueTitle forge.IssueTitle, prefix string, suffix string) string {
	title := prefix + issueTitle.Content + suffix
	// Using a single strings.NewReplacer
	replacer := strings.NewReplacer(
		" ", "-",
		"(", "",
		")", "",
	)
	return replacer.Replace(title)
}

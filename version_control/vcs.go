package version_control

import (
	"gokid/forge"
	"strings"
)

// VCS defines the interface for version control operations
type VCS interface {
	SyncTrunk(defaultBranch string) error
	ShowDiffSummary(branch string) error
	IsClean() (bool, error)
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

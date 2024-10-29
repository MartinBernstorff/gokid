package version_control

import (
	"gokid/forge"
	"strings"
)

// VCS defines the interface for version control operations
type VCS interface {
	NewChange(issue forge.Issue, defaultBranch string, migrateChanges bool, branchPrefix string, branchSuffix string) error
}

// NewChange implements VCS interface
func (g *BaseGit) NewChange(issue forge.Issue, defaultBranch string, migrateChanges bool, branchPrefix string, branchSuffix string) error {
	needsMigration := migrateChanges && !g.ops.isClean()

	if needsMigration {
		g.stash.Save()
	}

	branchTitle := branchTitle(issue, branchPrefix, branchSuffix)
	g.ops.fetch("origin")
	g.ops.branchFromOrigin(branchTitle, defaultBranch)

	if needsMigration {
		g.stash.Pop()
	}

	g.ops.emptyCommit(branchTitle)
	g.ops.push()
	return nil
}

// branchTitle creates a branch name from an issue title
func branchTitle(issue forge.Issue, prefix string, suffix string) string {
	title := prefix + issue.Title.Content + suffix
	return strings.ReplaceAll(title, " ", "-")
}

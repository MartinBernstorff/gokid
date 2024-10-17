package version_control

import (
	"fmt"
	"gokid/forge"
	"gokid/shell"
	"os/exec"
	"strings"
)

func branchTitle(issue forge.Issue, prefix string, suffix string) string {
	title := issue.Title.Content
	if prefix != "" {
		title += " " + prefix
	}
	if suffix != "" {
		title += " " + suffix
	}
	return strings.ReplaceAll(title, " ", "-")
}

func gitClean() bool {
	cmd := exec.Command("git", "status", "--porcelain")
	output, err := cmd.Output()
	return err != nil || strings.TrimSpace(string(output)) == ""
}

func NewChange(issue forge.Issue, defaultBranch string, migrateChanges bool, branchPrefix string, branchSuffix string) error {
	needsMigration := migrateChanges && !gitClean()

	if needsMigration {
		shell.Run("git stash")
	}

	branchTitle := branchTitle(issue, branchPrefix, branchSuffix)
	shell.Run("git fetch origin")
	shell.Run(fmt.Sprintf("git checkout -b %s --no-track origin/%s", branchTitle, defaultBranch))

	if needsMigration {
		shell.Run("git stash pop")
	}

	shell.Run(fmt.Sprintf("git commit --allow-empty -m '%s'", branchTitle))
	shell.Run("git push")
	return nil
}

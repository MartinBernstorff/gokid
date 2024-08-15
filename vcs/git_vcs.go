package vcs

import (
	"fmt"
	"gokid/forge"
	"gokid/shell"
	"os/exec"
	"strings"
)

func branchTitle(issue forge.Issue) string {
	title := issue.Title.Content
	return strings.ReplaceAll(title, " ", "_")
}

func gitClean() bool {
	cmd := exec.Command("git", "status", "--porcelain")
	output, err := cmd.Output()
	return err != nil || strings.TrimSpace(string(output)) == ""
}

func NewChange(issue forge.Issue, defaultBranch string, migrateChanges bool) error {
	needsMigration := migrateChanges && !gitClean()

	if needsMigration {
		shell.Run("git stash")
	}

	shell.Run("git fetch origin")
	shell.Run(fmt.Sprintf("git checkout -b %s --no-track origin/%s", branchTitle(issue), defaultBranch))

	if needsMigration {
		shell.Run("git stash pop")
	}

	shell.Run(fmt.Sprintf("git commit --allow-empty -m '%s'", branchTitle(issue)))
	shell.Run("git push")
	return nil
}

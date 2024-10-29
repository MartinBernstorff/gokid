package version_control

import (
	"fmt"
	"gokid/forge"
	"gokid/shell"
	"os/exec"
	"strings"
)

type GitVCS struct{}

func NewGitVCS() *GitVCS {
	return &GitVCS{}
}


func (g *GitVCS) IsClean() bool {
	cmd := exec.Command("git", "status", "--porcelain")
	output, err := cmd.Output()
	return err != nil || strings.TrimSpace(string(output)) == ""
}

func (g *GitVCS) StashChanges() {
	shell.Run("git stash")
}

func (g *GitVCS) PopStash() {
	shell.Run("git stash pop")
}

func (g *GitVCS) FetchOrigin() {
	shell.Run("git fetch origin")
}

func (g *GitVCS) CheckoutNewBranch(branchName string, baseBranch string) {
	shell.Run(fmt.Sprintf("git checkout -b %s --no-track origin/%s", branchName, baseBranch))
}

func (g *GitVCS) CreateEmptyCommit(message string) {
	shell.Run(fmt.Sprintf("git commit --allow-empty -m '%s'", message))
}

func (g *GitVCS) Push() {
	shell.Run("git push")
}

func (g *GitVCS) NewChange(issue forge.Issue, defaultBranch string, migrateChanges bool, branchPrefix string, branchSuffix string) error {
	base := NewBaseVCS(g)
	return base.NewChange(issue, defaultBranch, migrateChanges, branchPrefix, branchSuffix)
}

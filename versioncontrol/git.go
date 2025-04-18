package versioncontrol

import (
	"fmt"
	"gokid/forge"
	"gokid/shell"
	"os/exec"
	"strings"
)

type gitOperations interface {
	IsClean() (bool, error)
	CurrentBranch() (string, error)

	fetch(remote string) error
	branchFromOrigin(branchName string, defaultBranch string) error
	branchExists(branchName string) (bool, error)
	deleteBranch(branchName string) error
	switchBranch(branchName string) error
	emptyCommit(message string) error
	push(branchName forge.BranchName) error
}

type BaseGit struct {
	Ops   gitOperations
	Stash Stasher
}

type Git struct {
	BaseGit
	shell shell.Shell
}

func NewGit(s shell.Shell) *Git {
	g := &Git{
		shell: s,
	}
	g.Ops = g
	g.Stash = NewStash(s)
	return g
}

func (g *Git) ShowDiffSummary(branch string) error {
	_, err := g.shell.Run(fmt.Sprintf("git diff --stat %s", branch))
	return err
}

func (g *Git) branchExists(branchName string) (bool, error) {
	cmd := exec.Command("git", "branch", "--list", branchName)
	output, err := cmd.Output()

	if err != nil {
		return false, err
	}

	return strings.TrimSpace(string(output)) != "", err
}

func (g *Git) CurrentBranch() (string, error) {
	cmd := exec.Command("git", "branch", "--show-current")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

func (g *Git) switchBranch(branchName string) error {
	_, err := g.shell.RunQuietly(fmt.Sprintf("git checkout %s", branchName))
	return err
}

func (g *Git) deleteBranch(branchName string) error {
	_, err := g.shell.RunQuietly(fmt.Sprintf("git branch -D %s", branchName))
	return err
}

func (g *Git) IsClean() (bool, error) {
	cmd := exec.Command("git", "status", "--porcelain")
	output, err := cmd.Output()
	if err != nil {
		return false, err
	}
	return strings.TrimSpace(string(output)) == "", nil
}

func (g *Git) fetch(remote string) error {
	_, err := g.shell.RunQuietly(fmt.Sprintf("git fetch %s", remote))
	return err
}

func (g *Git) branchFromOrigin(branchName string, defaultBranch string) error {
	_, err := g.shell.RunQuietly(fmt.Sprintf("git checkout -b %s --no-track origin/%s", branchName, defaultBranch))
	return err
}

func (g *Git) emptyCommit(message string) error {
	_, err := g.shell.RunQuietly(fmt.Sprintf("git commit --allow-empty -m '%s'", message))
	return err
}

// `git push --set-upstream origin {branch_name}`
func (g *Git) push(branchName forge.BranchName) error {
	_, err := g.shell.RunQuietly(fmt.Sprintf("git push --set-upstream origin %s", branchName.String()))
	return err
}

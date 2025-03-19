package version_control

import (
	"fmt"
	"gokid/shell"
	"os/exec"
	"strings"
)

type gitOperations interface {
	IsClean() (bool, error)
	Fetch(remote string) error
	BranchFromOrigin(branchName string, defaultBranch string) error
	BranchExists(branchName string) (bool, error)
	DeleteBranch(branchName string) error
	SwitchBranch(branchName string) error
	EmptyCommit(message string) error
	Push() error
}

type BaseGit struct {
	Ops   gitOperations
	Stash Stasher
}

type Stasher interface {
	Save()
	Pop()
}

// Stash handles git stash operations
type Stash struct {
	shell shell.Shell
}

func NewStash(s shell.Shell) *Stash {
	return &Stash{
		shell: s,
	}
}

func (s *Stash) Save() {
	s.shell.Run("git stash")
}

func (s *Stash) Pop() {
	s.shell.Run("git stash pop")
}

// Git implements the VCS interface
type Git struct {
	BaseGit
	shell shell.Shell
}

// NewGit creates a new Git VCS instance
func NewGit(s shell.Shell) *Git {
	// p2: Check shell is in a Git repository
	g := &Git{
		shell: s,
	}
	g.Ops = g
	g.Stash = NewStash(s)
	return g
}

func (g *Git) ShowDiffSummary(branch string) error {
	g.shell.Run(fmt.Sprintf("git diff --stat %s", branch))
	return nil
}

func (g *Git) BranchExists(branchName string) (bool, error) {
	cmd := exec.Command("git", "branch", "--list", branchName)
	output, err := cmd.Output()

	if err != nil {
		return false, err
	}

	return strings.TrimSpace(string(output)) != "", err
}

func (g *Git) SwitchBranch(branchName string) error {
	_, err := g.shell.Run(fmt.Sprintf("git checkout %s", branchName))
	return err
}

func (g *Git) DeleteBranch(branchName string) error {
	_, err := g.shell.Run(fmt.Sprintf("git branch -D %s", branchName))
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

func (g *Git) Fetch(remote string) error {
	_, err := g.shell.Run(fmt.Sprintf("git fetch %s", remote))
	return err
}

func (g *Git) BranchFromOrigin(branchName string, defaultBranch string) error {
	_, err := g.shell.Run(fmt.Sprintf("git checkout -b %s --no-track origin/%s", branchName, defaultBranch))
	return err
}

func (g *Git) EmptyCommit(message string) error {
	_, err := g.shell.Run(fmt.Sprintf("git commit --allow-empty -m '%s'", message))
	return err

}

func (g *Git) Push() error {
	_, err := g.shell.Run("git push")
	return err
}

func (g *Git) SyncTrunk(defaultBranch string) error {
	g.shell.Run(fmt.Sprintf("git fetch origin %s", defaultBranch))
	g.shell.Run(fmt.Sprintf("git merge origin/%s", defaultBranch))
	return nil
}

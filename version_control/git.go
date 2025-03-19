package version_control

import (
	"fmt"
	"gokid/forge"
	"gokid/shell"
	"os/exec"
	"strings"
)

// p2: Change these methods to return errors, rather than panic on error
type gitOperations interface {
	IsClean() bool
	Fetch(remote string)
	BranchFromOrigin(branchName string, defaultBranch string)
	BranchExists(branchName string) bool
	EmptyCommit(message string)
	Push()
}

// BaseGit implements common git functionality
type BaseGit struct {
	Ops   gitOperations
	Stash Stasher
}

// NewChange implements VCS interface
func (g *BaseGit) NewChange(issue forge.Issue, defaultBranch string, migrateChanges bool, branchPrefix string, branchSuffix string) error {
	needsMigration := migrateChanges && !g.Ops.IsClean()
	if needsMigration {
		g.Stash.Save()
	}

	branchTitle := BranchTitle(issue.Title, branchPrefix, branchSuffix)
	g.Ops.Fetch("origin")
	g.Ops.BranchFromOrigin(branchTitle, defaultBranch)

	if needsMigration {
		g.Stash.Pop()
	}

	g.Ops.EmptyCommit("Initial commit")
	g.Ops.Push()
	return nil
}

// Stasher defines the interface for stash operations
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

func (g *Git) BranchExists(branchName string) bool {
	cmd := exec.Command("git", "branch", "--list", branchName)
	output, err := cmd.Output()
	return err == nil && strings.TrimSpace(string(output)) != ""
}

func (g *Git) IsClean() bool {
	cmd := exec.Command("git", "status", "--porcelain")
	output, err := cmd.Output()
	return err != nil || strings.TrimSpace(string(output)) == ""
}

func (g *Git) Fetch(remote string) {
	g.shell.Run(fmt.Sprintf("git fetch %s", remote))
}

func (g *Git) BranchFromOrigin(branchName string, defaultBranch string) {
	g.shell.Run(fmt.Sprintf("git checkout -b %s --no-track origin/%s", branchName, defaultBranch))
}

func (g *Git) EmptyCommit(message string) {
	g.shell.Run(fmt.Sprintf("git commit --allow-empty -m '%s'", message))
}

func (g *Git) Push() {
	g.shell.Run("git push")
}

func (g *Git) SyncTrunk(defaultBranch string) error {
	g.shell.Run(fmt.Sprintf("git fetch origin %s", defaultBranch))
	g.shell.Run(fmt.Sprintf("git merge origin/%s", defaultBranch))
	return nil
}

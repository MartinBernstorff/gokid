package version_control

import (
	"fmt"
	"gokid/forge"
	"gokid/shell"
	"os/exec"
	"strings"
)

// gitOperations defines the interface for git operations that can be implemented differently by Git and GitStub
type gitOperations interface {
	isClean() bool
	fetch(remote string)
	branchFromOrigin(branchName string, defaultBranch string)
	emptyCommit(message string)
	push()
}

// BaseGit implements common git functionality
type BaseGit struct {
	ops   gitOperations
	stash Stasher
}

// NewChange implements VCS interface
func (g *BaseGit) NewChange(issue forge.Issue, defaultBranch string, migrateChanges bool, branchPrefix string, branchSuffix string) error {
	needsMigration := migrateChanges && !g.ops.isClean()
	if needsMigration {
		g.stash.Save()
	}

	branchTitle := branchTitle(issue.Title, branchPrefix, branchSuffix)
	g.ops.fetch("origin")
	g.ops.branchFromOrigin(branchTitle, defaultBranch)

	if needsMigration {
		g.stash.Pop()
	}

	g.ops.emptyCommit("Initial commit")
	g.ops.push()
	return nil
}

// Stasher defines the interface for stash operations
type Stasher interface {
	Save()
	Pop()
}

// Stash handles git stash operations
type Stash struct{}

func NewStash() *Stash {
	return &Stash{}
}

func (s *Stash) Save() {
	myShell := shell.New()
	myShell.Run("git stash")
}

func (s *Stash) Pop() {
	myShell := shell.New()
	myShell.Run("git stash pop")
}

// Git implements the VCS interface
type Git struct {
	BaseGit
}

// NewGit creates a new Git VCS instance
func NewGit() *Git {
	g := &Git{}
	g.ops = g
	g.stash = NewStash()
	return g
}

func (g *Git) isClean() bool {
	cmd := exec.Command("git", "status", "--porcelain")
	output, err := cmd.Output()
	return err != nil || strings.TrimSpace(string(output)) == ""
}

func (g *Git) fetch(remote string) {
	myShell := shell.New()
	myShell.Run(fmt.Sprintf("git fetch %s", remote))
}

func (g *Git) branchFromOrigin(branchName string, defaultBranch string) {
	myShell := shell.New()
	myShell.Run(fmt.Sprintf("git checkout -b %s --no-track origin/%s", branchName, defaultBranch))
}

func (g *Git) emptyCommit(message string) {
	myShell := shell.New()
	myShell.Run(fmt.Sprintf("git commit --allow-empty -m '%s'", message))
}

func (g *Git) push() {
	myShell := shell.New()
	myShell.Run("git push")
}

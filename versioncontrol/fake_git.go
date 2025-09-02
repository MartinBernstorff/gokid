package versioncontrol

import (
	"fmt"
	"gokid/forge"

	"github.com/samber/lo"
)

// FakeStash maintains a simple stash counter and manages dirty state
type FakeStash struct {
	stashCount int
	git        *FakeGit // Reference to parent git to manage dirty state
}

func NewFakeStash(git *FakeGit) *FakeStash {
	return &FakeStash{
		git: git,
	}
}

func (s *FakeStash) Save() error {
	s.stashCount++
	s.git.isDirty = false // Stashing makes working directory clean
	return nil
}

func (s *FakeStash) Pop() error {
	if s.stashCount > 0 {
		s.stashCount--
		s.git.isDirty = true // Popping makes working directory dirty again
	}
	return nil
}

// Commit represents a git commit with minimal information
type Commit struct {
	Title string
	Empty bool
}

// FakeGit simulates minimal git repository state
type FakeGit struct {
	BaseGit

	// Repository state
	currentBranchField string
	originBranch       string
	branches           []string
	isDirty            bool
	commits            []Commit
	lastPush           Commit
	isFetched          bool
	TrunkSynced        bool
	DiffSummaryCalls   int
}

func NewFakeGit() *FakeGit {
	g := &FakeGit{
		originBranch:     "main", // default origin branch
		commits:          make([]Commit, 0),
		isFetched:        false,
		TrunkSynced:      false,
		DiffSummaryCalls: 0,
	}
	g.Ops = g
	g.Stash = NewFakeStash(g) // Pass git reference to stash
	return g
}

func (g *FakeGit) branchExists(branchName string) (bool, error) {
	for _, branch := range g.branches {
		if branch == branchName {
			return true, nil
		}
	}
	return false, nil
}

func (g *FakeGit) deleteBranch(branchName string) error {
	filteredBranches := lo.Filter(g.branches, func(branch string, _ int) bool {
		return branch != branchName
	})

	if len(filteredBranches) == len(g.branches) {
		return fmt.Errorf("branch %s not found", branchName)
	}

	g.branches = filteredBranches
	g.currentBranchField = ""
	return nil
}

func (g *FakeGit) CurrentBranch() (string, error) {
	return g.currentBranchField, nil
}

func (g *FakeGit) switchBranch(branchName string) error {
	exists, err := g.branchExists(branchName)

	if err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("branch %s not found", branchName)
	}

	g.currentBranchField = branchName
	return nil
}

func (g *FakeGit) OriginBranch() string {
	return g.originBranch
}

func (g *FakeGit) IsDirty() bool {
	return g.isDirty
}

func (g *FakeGit) StashCount() int {
	return g.Stash.(*FakeStash).stashCount
}

func (g *FakeGit) Commits() []Commit {
	return g.commits
}

// Implementation of gitOperations interface
func (g *FakeGit) SetDirty(isDirty bool) {
	g.isDirty = isDirty
}

func (g *FakeGit) IsClean() (bool, error) {
	return !g.isDirty, nil
}

func (g *FakeGit) fetch(_ string, branch string) error {
	g.isFetched = true
	return nil
}

func (g *FakeGit) branchFromOrigin(branchName string, origin string) error {
	g.branches = append(g.branches, branchName)
	g.currentBranchField = branchName
	g.originBranch = origin

	return nil
}

func (g *FakeGit) emptyCommit(message string) error {
	g.commits = append(g.commits, Commit{
		Title: message,
		Empty: true,
	})
	return nil
}

func (g *FakeGit) push(_ forge.BranchName) error {
	g.lastPush = g.commits[len(g.commits)-1]
	return nil
}

func (g *FakeGit) AddCommit(title string, empty bool) {
	g.commits = append(g.commits, Commit{
		Title: title,
		Empty: empty,
	})
}

func (g *FakeGit) SyncTrunk(_ string) error {
	g.TrunkSynced = true
	return nil
}

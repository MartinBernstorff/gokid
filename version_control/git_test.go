package version_control

import (
	"gokid/forge"
	"testing"
)

// FakeVCS implements VCS interface for testing
type FakeVCS struct {
	clean          bool
	currentBranch  string
	stash          []string
	commitMessages []string
	pushed         bool
	fetchCalled    bool
	stashOps       int  // Track total stash operations (stash + pop)
}

func NewFakeVCS() *FakeVCS {
	return &FakeVCS{
		clean:          true,
		stash:          make([]string, 0),
		commitMessages: make([]string, 0),
	}
}

func (f *FakeVCS) IsClean() bool {
	return f.clean
}

func (f *FakeVCS) StashChanges() {
	f.stash = append(f.stash, "changes")
	f.clean = true
	f.stashOps++
}

func (f *FakeVCS) PopStash() {
	if len(f.stash) > 0 {
		f.stash = f.stash[:len(f.stash)-1]
		f.clean = false
		f.stashOps++
	}
}

func (f *FakeVCS) FetchOrigin() {
	f.fetchCalled = true
}

func (f *FakeVCS) CheckoutNewBranch(branchName string, baseBranch string) {
	f.currentBranch = branchName
}

func (f *FakeVCS) CreateEmptyCommit(message string) {
	f.commitMessages = append(f.commitMessages, message)
}

func (f *FakeVCS) Push() {
	f.pushed = true
}

func (f *FakeVCS) NewChange(issue forge.Issue, defaultBranch string, migrateChanges bool, branchPrefix string, branchSuffix string) error {
	base := NewBaseVCS(f)
	return base.NewChange(issue, defaultBranch, migrateChanges, branchPrefix, branchSuffix)
}

func TestNewChangeWithFakeVCS(t *testing.T) {
	tests := []struct {
		name           string
		issue          forge.Issue
		defaultBranch  string
		migrateChanges bool
		branchPrefix   string
		branchSuffix   string
		initialClean   bool
		wantBranch     string
		wantCommit     string
		wantDirty      bool
		wantStashOps   int
	}{
		{
			name: "Clean repository - no migration needed",
			issue: forge.Issue{
				Title: forge.IssueTitle{Content: "test feature"},
			},
			defaultBranch:  "main",
			migrateChanges: true,
			branchPrefix:   "feature/",
			branchSuffix:   "",
			initialClean:   true,
			wantBranch:     "feature/test-feature",
			wantCommit:     "feature/test feature",
			wantDirty:      false,
			wantStashOps:   0,
		},
		{
			name: "Dirty repository with migration",
			issue: forge.Issue{
				Title: forge.IssueTitle{Content: "test feature"},
			},
			defaultBranch:  "main",
			migrateChanges: true,
			branchPrefix:   "feature/",
			branchSuffix:   "",
			initialClean:   false,
			wantBranch:     "feature/test-feature",
			wantCommit:     "feature/test feature",
			wantDirty:      true,
			wantStashOps:   2, // stash + pop
		},
		{
			name: "Dirty repository without migration",
			issue: forge.Issue{
				Title: forge.IssueTitle{Content: "test feature"},
			},
			defaultBranch:  "main",
			migrateChanges: false,
			branchPrefix:   "feature/",
			branchSuffix:   "",
			initialClean:   false,
			wantBranch:     "feature/test-feature",
			wantCommit:     "feature/test feature",
			wantDirty:      false,
			wantStashOps:   0,
		},
		{
			name: "With branch suffix",
			issue: forge.Issue{
				Title: forge.IssueTitle{Content: "test feature"},
			},
			defaultBranch:  "main",
			migrateChanges: true,
			branchPrefix:   "feature/",
			branchSuffix:   "-dev",
			initialClean:   true,
			wantBranch:     "feature/test-feature-dev",
			wantCommit:     "feature/test feature-dev",
			wantDirty:      false,
			wantStashOps:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vcs := NewFakeVCS()
			vcs.clean = tt.initialClean

			err := vcs.NewChange(tt.issue, tt.defaultBranch, tt.migrateChanges, tt.branchPrefix, tt.branchSuffix)
			if err != nil {
				t.Errorf("NewChange() error = %v", err)
				return
			}

			if !vcs.fetchCalled {
				t.Error("NewChange() did not call FetchOrigin")
			}

			if vcs.currentBranch != tt.wantBranch {
				t.Errorf("NewChange() branch = %v, want %v", vcs.currentBranch, tt.wantBranch)
			}

			if len(vcs.commitMessages) == 0 || vcs.commitMessages[0] != tt.wantCommit {
				t.Errorf("NewChange() commit = %v, want %v", vcs.commitMessages, tt.wantCommit)
			}

			if vcs.clean == tt.wantDirty {
				t.Errorf("NewChange() clean = %v, want dirty = %v", vcs.clean, tt.wantDirty)
			}

			if !vcs.pushed {
				t.Error("NewChange() changes were not pushed")
			}

			actualStashOps := vcs.stashOps
			if actualStashOps != tt.wantStashOps {
				t.Errorf("NewChange() stash operations = %v, want %v", actualStashOps, tt.wantStashOps)
			}
		})
	}
}

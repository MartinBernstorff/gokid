package version_control

import (
	"gokid/forge"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewChange(t *testing.T) {
	t.Run("creates new branch with empty commit", func(t *testing.T) {
		// Arrange
		git := NewGitStub()
		issue := forge.Issue{Title: forge.IssueTitle{Content: "test issue"}}

		// Act
		err := git.NewChange(issue, "main", false, "prefix-", "-suffix")

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, "prefix-test-issue-suffix", git.CurrentBranch())
		assert.Equal(t, "main", git.OriginBranch())

		commits := git.Commits()
		assert.Len(t, commits, 1)
		assert.Equal(t, "prefix-test-issue-suffix", commits[0].Title)
		assert.True(t, commits[0].Empty)
		assert.True(t, git.remoteIsUpdated())
	})

	t.Run("handles dirty working directory with migration", func(t *testing.T) {
		// Arrange
		git := NewGitStub()
		git.SetDirty(true)
		issue := forge.Issue{Title: forge.IssueTitle{Content: "test issue"}}

		// Act
		err := git.NewChange(issue, "main", true, "prefix-", "-suffix")

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 0, git.StashCount()) // Stashed and popped
		assert.Equal(t, "prefix-test-issue-suffix", git.CurrentBranch())

		commits := git.Commits()
		assert.Len(t, commits, 1)
		assert.True(t, git.remoteIsUpdated())
	})

	t.Run("skips stash when working directory is clean", func(t *testing.T) {
		// Arrange
		git := NewGitStub()
		git.SetDirty(false)
		issue := forge.Issue{Title: forge.IssueTitle{Content: "test issue"}}

		// Act
		err := git.NewChange(issue, "main", true, "prefix-", "-suffix")

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 0, git.StashCount()) // No stash operations
		assert.Equal(t, "prefix-test-issue-suffix", git.CurrentBranch())
	})

	t.Run("skips stash when migration is disabled", func(t *testing.T) {
		// Arrange
		git := NewGitStub()
		git.SetDirty(true)
		issue := forge.Issue{Title: forge.IssueTitle{Content: "test issue"}}

		// Act
		err := git.NewChange(issue, "main", false, "prefix-", "-suffix")

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 0, git.StashCount()) // No stash operations
		assert.Equal(t, "prefix-test-issue-suffix", git.CurrentBranch())
	})

	t.Run("preserves existing commits when creating new branch", func(t *testing.T) {
		// Arrange
		git := NewGitStub()
		git.AddCommit("initial commit", false)
		issue := forge.Issue{Title: forge.IssueTitle{Content: "test issue"}}

		// Act
		err := git.NewChange(issue, "main", false, "prefix-", "-suffix")

		// Assert
		assert.NoError(t, err)
		commits := git.Commits()
		assert.Len(t, commits, 2)
		assert.Equal(t, "initial commit", commits[0].Title)
		assert.False(t, commits[0].Empty)
		assert.Equal(t, "prefix-test-issue-suffix", commits[1].Title)
		assert.True(t, commits[1].Empty)
	})
}

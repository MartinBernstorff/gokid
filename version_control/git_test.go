package version_control

import (
	"gokid/forge"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewChange(t *testing.T) {
	issue := forge.Issue{Title: forge.IssueTitle{Content: "test issue"}}

	t.Run("creates new branch with empty commit", func(t *testing.T) {
		git := NewFakeGit()
		expectedBranchName := branchTitle(issue, "prefix-", "-suffix")

		git.NewChange(issue, "main", false, "prefix-", "-suffix")

		// Remote state
		assert.True(t, git.remoteIsUpdated())
		assert.Equal(t, "main", git.OriginBranch())
		assert.True(t, git.isFetched)

		// Local state
		commits := git.Commits()
		assert.Len(t, commits, 1)
		assert.True(t, commits[0].Empty)

		assert.Equal(t, expectedBranchName, git.CurrentBranch())
		assert.True(t, git.isClean())
	})

	t.Run("migrates dirty changes", func(t *testing.T) {
		git := NewFakeGit()
		git.SetDirty(true)

		git.NewChange(issue, "main", true, "", "")

		assert.Equal(t, 0, git.StashCount()) // Stashed and popped
		assert.False(t, git.isClean())       // Stash is applied
		assert.Len(t, git.Commits(), 1)      // Has only the initial commit
	})
}

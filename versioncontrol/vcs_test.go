package versioncontrol

import (
	"gokid/forge"
	"testing"
)

func TestBranchTitle(t *testing.T) {
	tests := []struct {
		name          string
		issueTitle    forge.IssueTitle
		expectedTitle string
	}{
		{
			name: "basic title",
			issueTitle: forge.IssueTitle{
				Content: "simple issue",
			},
			expectedTitle: "simple-issue",
		},
		{
			name: "title with prefix",
			issueTitle: forge.IssueTitle{
				Content: "add feature",
			},
			expectedTitle: "add-feature",
		},
		{
			name: "title with parentheses",
			issueTitle: forge.IssueTitle{
				Content: "fix bug (urgent)",
			},
			expectedTitle: "fix-bug-urgent",
		},
		{
			name: "title with prefix and suffix",
			issueTitle: forge.IssueTitle{
				Content: "add feature",
			},
			expectedTitle: "add-feature",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.issueTitle.ToBranchName()

			if got.String() != tt.expectedTitle {
				t.Errorf("got %v, want %v", got, tt.expectedTitle)
			}
		})
	}
}

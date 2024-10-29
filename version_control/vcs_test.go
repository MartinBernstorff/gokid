package version_control

import (
	"gokid/forge"
	"testing"
)

func TestBranchTitle(t *testing.T) {
	tests := []struct {
		name          string
		issueTitle    forge.IssueTitle
		prefix        string
		suffix        string
		expectedTitle string
	}{
		{
			name: "basic title",
			issueTitle: forge.IssueTitle{
				Content: "simple issue",
			},
			prefix:        "",
			suffix:        "",
			expectedTitle: "simple-issue",
		},
		{
			name: "title with parentheses",
			issueTitle: forge.IssueTitle{
				Content: "fix bug (urgent)",
			},
			prefix:        "",
			suffix:        "",
			expectedTitle: "fix-bug-urgent",
		},
		{
			name: "title with prefix and suffix",
			issueTitle: forge.IssueTitle{
				Content: "add feature",
			},
			prefix:        "feature/",
			suffix:        "-123",
			expectedTitle: "feature/add-feature-123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := branchTitle(tt.issueTitle, tt.prefix, tt.suffix)
			if got != tt.expectedTitle {
				t.Errorf("branchTitle() = %v, want %v", got, tt.expectedTitle)
			}
		})
	}
}

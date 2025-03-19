package forge

import (
	"reflect"
	"testing"
)

func TestParseIssueTitle(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected IssueTitle
	}{
		{
			name:     "Normal case with prefix and description",
			input:    "BUG: Application crashes on startup",
			expected: IssueTitle{Prefix: "BUG", Content: "Application crashes on startup"},
		},
		{
			name:     "No prefix, only description",
			input:    "Application crashes on startup",
			expected: IssueTitle{Prefix: "", Content: "Application crashes on startup"},
		},
		{
			name:     "Multiple colons in description",
			input:    "TEST: Verify login: success and failure cases",
			expected: IssueTitle{Prefix: "TEST", Content: "Verify login: success and failure cases"},
		},
		{
			name:     "Empty input",
			input:    "",
			expected: IssueTitle{Prefix: "", Content: ""},
		},
		{
			name:     "Only prefix, no description",
			input:    "PREFIX:",
			expected: IssueTitle{Prefix: "PREFIX", Content: ""},
		},
		{
			name:     "Monorepo-style",
			input:    "fix(component): context",
			expected: IssueTitle{Prefix: "fix(component)", Content: "context"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ParseIssueTitle(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ParseIssueTitle(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

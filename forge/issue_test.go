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
			name:     "Prefix with parenthesis",
			input:    "FEATURE(UI): Add dark mode",
			expected: IssueTitle{Prefix: "FEATURE", Content: "Add dark mode"},
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
			input:    "TODO:",
			expected: IssueTitle{Prefix: "TODO", Content: ""},
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

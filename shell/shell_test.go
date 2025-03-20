package shell

import (
	"strings"
	"testing"
)

func TestError(t *testing.T) {
	shell := New()
	output, err := shell.Run("ls /nonexistent")

	if err == nil {
		t.Fatal("expected error for nonexistent directory")
	}

	expectedErr := "ls: /nonexistent: No such file or directory"
	if !strings.Contains(output, expectedErr) {
		t.Errorf("got output %q, want output containing %q", output, expectedErr)
	}
}

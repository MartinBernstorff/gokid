package shell

import "testing"

func TestError(t *testing.T) {
	shell := New()
	_, err := shell.Run("ls /nonexistent")
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

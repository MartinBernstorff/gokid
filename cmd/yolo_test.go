package cmd

import (
	"gokid/forge"
	"gokid/version_control"
	"testing"
)

func TestYolo(t *testing.T) {
	tests := []struct {
		name          string
		draft         bool
		mergeStrategy string
		userConfirmed bool
		// Only set the fields we want to check in each test
		wantStrategy string
		wantForced   bool
		wantDraft    bool
	}{
		{
			name:          "unconfirmed yolo aborts without calling merge",
			userConfirmed: false,
			wantForced:    false,
		},
		{
			name:          "yolo preserves merge strategy",
			mergeStrategy: "squash",
			userConfirmed: true,
			wantStrategy:  "squash",
		},
		{
			name:          "yolo forces merge",
			userConfirmed: true,
			wantForced:    true,
		},
		{
			name:          "yolo with draft PR marks ready",
			draft:         true,
			userConfirmed: true,
			wantDraft:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup fake forge and shell
			fakeForge := forge.NewFakeForge()
			merger := NewMerger(fakeForge, version_control.NewFakeGit())
			yoloer := NewYoloer(merger)

			// Run yolo command
			yoloer.yolo(tt.draft, tt.mergeStrategy, tt.userConfirmed)

			// Check if merge was called when it shouldn't have been
			if !tt.userConfirmed && fakeForge.LastMergeStrategy != "" {
				t.Error("merge was called when it should have been aborted")
				return
			}

			if tt.wantStrategy != "" && fakeForge.LastMergeStrategy != tt.wantStrategy {
				t.Errorf("merge strategy = %v, want %v", fakeForge.LastMergeStrategy, tt.mergeStrategy)
			}

			if tt.wantForced {
				if fakeForge.LastAutoMerge {
					t.Error("auto merge should be disabled")
				}
				if !fakeForge.LastForceMerge {
					t.Error("force merge should be enabled")
				}
			}
			if tt.wantDraft && fakeForge.WasMarkedReady != tt.wantDraft {
				t.Errorf("marked ready = %v, want %v", fakeForge.WasMarkedReady, tt.wantDraft)
			}
		})
	}
}

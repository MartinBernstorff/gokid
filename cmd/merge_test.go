package cmd

import (
	"gokid/forge"
	"gokid/versioncontrol"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name             string
		preMergeCmd      string
		autoMerge        bool
		forceMerge       bool
		draft            bool
		mergeStrategy    string
		wantStrategy     string
		wantAutoMerge    bool
		wantForceMerge   bool
		wantReady        bool
		syncTrunkOnMerge bool
		trunk            string
		wantSyncCalled   bool
	}{
		{
			name:          "merge strategy is passed to forge",
			mergeStrategy: "squash",
			wantStrategy:  "squash",
		},
		{
			name:          "automerge is passed to forge",
			autoMerge:     true,
			wantAutoMerge: true,
		},
		{
			name:           "force merge is passed to forge",
			forceMerge:     true,
			wantForceMerge: true,
		},
		{
			name:      "draft PR is marked ready before merge",
			draft:     true,
			wantReady: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup fake forge and VCS
			fakeForge := forge.NewFakeForge()
			fakeVCS := versioncontrol.NewFakeGit()
			merger := NewMerger(fakeForge, fakeVCS)

			// Run merge command
			merger.merge(tt.preMergeCmd, tt.autoMerge, tt.forceMerge, tt.draft, tt.mergeStrategy)

			// Verify forge calls
			if fakeForge.LastMergeStrategy != tt.wantStrategy {
				t.Errorf("merge strategy = %v, want %v", fakeForge.LastMergeStrategy, tt.wantStrategy)
			}

			if fakeForge.LastAutoMerge != tt.wantAutoMerge {
				t.Errorf("auto merge = %v, want %v", fakeForge.LastAutoMerge, tt.wantAutoMerge)
			}

			if fakeForge.LastForceMerge != tt.wantForceMerge {
				t.Errorf("force merge = %v, want %v", fakeForge.LastForceMerge, tt.wantForceMerge)
			}

			if fakeForge.WasMarkedReady != tt.wantReady {
				t.Errorf("marked ready = %v, want %v", fakeForge.WasMarkedReady, tt.wantReady)
			}
		})
	}
}

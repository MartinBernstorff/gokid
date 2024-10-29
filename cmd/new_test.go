package cmd

import (
	"gokid/config"
	"gokid/forge"
	"gokid/version_control"
	"testing"
)

func TestNewChange(t *testing.T) {
	tests := []struct {
		name       string
		config     *config.GokidConfig
		inputTitle string
		wantPR     forge.PullRequest
		wantErr    bool
	}{
		{
			name: "Create draft PR with prefix",
			config: &config.GokidConfig{
				Trunk: "main",
				Draft: true,
			},
			wantPR: forge.PullRequest{
				Title: "FEAT: New feature",
				Base:  "main",
				Draft: true,
			},
		},
		{
			name: "Create PR without prefix",
			config: &config.GokidConfig{
				Trunk: "master",
				Draft: false,
			},
			wantPR: forge.PullRequest{
				Title: "Simple change",
				Base:  "master",
				Draft: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fakeForge := forge.NewFakeForge()

			err := newChange(fakeForge, tt.config, tt.inputTitle, version_control.NewFakeGit())

			if (err != nil) != tt.wantErr {
				t.Errorf("newChange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			got := fakeForge.LastCreatedPR
			if got.Base != tt.wantPR.Base {
				t.Errorf("PR Base = %v, want %v", got.Base, tt.wantPR.Base)
			}
			if got.Draft != tt.wantPR.Draft {
				t.Errorf("PR Draft = %v, want %v", got.Draft, tt.wantPR.Draft)
			}
		})
	}
}

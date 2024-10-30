/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"gokid/config"
	"gokid/forge"
	"gokid/shell"
	"gokid/version_control"

	"github.com/spf13/cobra"
)

type Merger struct {
	forge forge.Forge
	shell shell.Shell
	vcs   version_control.VCS
}

func NewMerger(f forge.Forge, v version_control.VCS) *Merger {
	return &Merger{
		forge: f,
		vcs:   v,
	}
}

func (m *Merger) merge(preMergeCommand string, autoMerge bool, forceMerge bool, draft bool, mergeStrategy string, trunk string, syncTrunkOnMerge bool) {
	if syncTrunkOnMerge {
		if err := m.vcs.SyncTrunk(trunk); err != nil {
			fmt.Println("Error syncing trunk:", err)
			return
		}
	}

	// Execute pre-merge command if set
	if preMergeCommand != "" {
		fmt.Println("Running pre-merge command:", preMergeCommand)
		m.shell.Run(preMergeCommand)
		fmt.Println("Pre-merge command completed")
	}

	if draft {
		if err := m.forge.MarkPullRequestReady(); err != nil {
			fmt.Println("Error marking PR as ready:", err)
			return
		}
	}

	if err := m.forge.MergePullRequest(mergeStrategy, autoMerge, forceMerge); err != nil {
		fmt.Println("Error merging PR:", err)
		return
	}
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "merge",
		Short: "Merge a change",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			cfg := config.Load(config.DefaultFileName)
			shell := shell.New()
			git := version_control.NewGit(shell)
			merger := NewMerger(forge.NewGitHub(shell), git)
			merger.merge(cfg.PreMergeCommand, cfg.AutoMerge, cfg.ForceMerge, cfg.Draft, cfg.MergeStrategy, cfg.Trunk, cfg.SyncTrunkOnMerge)
		},
		Aliases: []string{"m"},
	})
}

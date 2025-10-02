/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"gokid/config"
	"gokid/forge"
	"gokid/shell"
	"gokid/versioncontrol"

	"github.com/spf13/cobra"
)

type Merger struct {
	forge forge.Forge
	vcs   versioncontrol.VCS
}

func NewMerger(f forge.Forge, v versioncontrol.VCS) *Merger {
	return &Merger{
		forge: f,
		vcs:   v,
	}
}

func (m *Merger) merge(preMergeCommand string, autoMerge bool, forceMerge bool, draft bool, mergeStrategy string, postMergeCommand string) {
	// Execute pre-merge command if set
	if preMergeCommand != "" {
		fmt.Println("Running pre-merge command:", preMergeCommand)
		shell := shell.New()
		_, err := shell.Run(preMergeCommand)
		if err != nil {
			fmt.Println("Error running pre-merge command:", err)
			return
		}

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
	fmt.Printf("Merge initiated - Strategy: %s, AutoMerge: %t, ForceMerge: %t",
		mergeStrategy, autoMerge, forceMerge)

	// Execute post-merge command if set
	if preMergeCommand != "" {
		fmt.Println("Running post-merge command:", postMergeCommand)
		shell := shell.New()
		_, err := shell.Run(preMergeCommand)
		if err != nil {
			fmt.Println("Error running pre-merge command:", err)
			return
		}

		fmt.Println("Post-merge command completed")
	}
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "merge",
		Short: "Merge a change",
		Long:  "",
		Run: func(_ *cobra.Command, _ []string) {
			cfg := config.Load(config.DefaultFileName)
			shell := shell.New()
			git := versioncontrol.NewGit(shell)
			merger := NewMerger(forge.NewGitHub(shell), git)
			merger.merge(cfg.PreMergeCommand, cfg.AutoMerge, cfg.ForceMerge, cfg.Draft, cfg.MergeStrategy, cfg.PostMergeCommand)
		},
		Aliases: []string{"m"},
	})
}

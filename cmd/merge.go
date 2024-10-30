/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"gokid/config"
	"gokid/forge"
	"gokid/shell"

	"github.com/spf13/cobra"
)

type Merger struct {
	forge forge.Forge
	shell shell.Shell
}

func NewMerger(f forge.Forge, s shell.Shell) *Merger {
	return &Merger{
		forge: f,
		shell: s,
	}
}

func (m *Merger) merge(preMergeCommand string, autoMerge bool, forceMerge bool, draft bool, mergeStrategy string) {
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
			merger := NewMerger(forge.NewGitHub(), shell.New())
			merger.merge(cfg.PreMergeCommand, cfg.AutoMerge, cfg.ForceMerge, cfg.Draft, cfg.MergeStrategy)
		},
		Aliases: []string{"m"},
	})
}

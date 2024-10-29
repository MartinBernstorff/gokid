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

func merge(cfg config.GokidConfig) {
	// Execute pre-merge command if set
	if cfg.PreMergeCommand != "" {
		fmt.Println("Running pre-merge command:", cfg.PreMergeCommand)
		shell.Run(cfg.PreMergeCommand)
		fmt.Println("Pre-merge command completed")
	}

	forge := forge.NewGitHub()

	if cfg.Draft {
		if err := forge.MarkPullRequestReady(); err != nil {
			fmt.Println("Error marking PR as ready:", err)
			return
		}
	}

	if err := forge.MergePullRequest(cfg.MergeStrategy, cfg.AutoMerge, cfg.ForceMerge); err != nil {
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
			merge(config.Load(config.DefaultFileName))
		},
		Aliases: []string{"m"},
	})
}

/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"gokid/config"
	"gokid/shell"
	"gokid/versioncontrol"

	"github.com/spf13/cobra"
)

var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Show diff against trunk",
	Long:  "Show the git diff between the current branch and trunk (main/master)",
	Run: func(cmd *cobra.Command, _ []string) {
		cfg := config.Load(config.DefaultFileName)
		shell := shell.New()
		git := versioncontrol.NewGit(shell)
		err := git.ShowDiffSummary(cfg.Trunk)
		if err != nil {
			cmd.PrintErrf("Error showing diff: %v\n", err)
		}
	},
	Aliases: []string{"d"},
}

func init() {
	rootCmd.AddCommand(diffCmd)
}

/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"gokid/config"
	"gokid/shell"

	"github.com/spf13/cobra"
)

func merge() {
	cfg := config.Load(config.DefaultFileName)

	if cfg.Draft {
		shell.Run("gh pr ready")
	}

	cmd := "gh pr merge"
	cmd += fmt.Sprintf(" --%s", cfg.MergeStrategy)

	if cfg.AutoMerge {
		cmd += " --auto"
	}

	shell.Run(cmd)
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "merge",
		Short: "Merge a change",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			merge()
		},
		Aliases: []string{"m"},
	})
}

/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"gokid/config"
	"gokid/shell"
	"slices"

	"github.com/spf13/cobra"
)

func merge() {
	cfg := config.Init()

	cmd := "gh pr merge"

	allowedStrategies := []string{"squash", "rebase", "merge"}

	if !slices.Contains(allowedStrategies, cfg.MergeStrategy) {
		fmt.Printf("Merge strategy is not allowed, allowed are: %s", allowedStrategies)
	} else {
		cmd += fmt.Sprintf(" --%s", cfg.MergeStrategy)
	}

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

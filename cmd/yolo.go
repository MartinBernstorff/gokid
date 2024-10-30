/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"gokid/config"

	"github.com/spf13/cobra"
)

func yolo(cfg config.GokidConfig) {
	// Override config for yolo mode
	cfg.ForceMerge = true
	cfg.PreMergeCommand = ""
	cfg.Yolo = true

	// Use the existing merge function with our modified config
	merge(cfg)
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "yolo",
		Short: "Merge a change without running checks",
		Long:  "YOLO mode merges changes without running pre-merge checks. Use with caution!",
		Run: func(cmd *cobra.Command, args []string) {
			yolo(config.Load(config.DefaultFileName))
		},
		Aliases: []string{"y"},
	})
}

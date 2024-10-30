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

type Yoloer struct {
	merger *Merger
}

func NewYoloer(merger *Merger) *Yoloer {
	return &Yoloer{
		merger: merger,
	}
}

func (y *Yoloer) yolo(draft bool, mergeStrategy string, confirmed bool) {
	if !confirmed {
		fmt.Println("Aborted.")
		return
	}
	y.merger.merge("", false, true, draft, mergeStrategy)
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "yolo",
		Short: "Merge a change without running checks",
		Long:  "YOLO mode merges changes without running pre-merge checks. Use with caution!",
		Run: func(cmd *cobra.Command, args []string) {
			cfg := config.Load(config.DefaultFileName)

			fmt.Println("🚀 YOLO mode enabled - merging without checks!")
			fmt.Println("Are you sure you want to merge? (y/n)")
			var confirm string
			fmt.Scanln(&confirm)

			merger := NewMerger(forge.NewGitHub(shell.New()))

			yoloer := NewYoloer(merger)
			yoloer.yolo(cfg.Draft, cfg.MergeStrategy, confirm == "y")
		},
	})
}

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

type Yoloer struct {
	merger *Merger
}

func NewYoloer(merger *Merger) *Yoloer {
	return &Yoloer{
		merger: merger,
	}
}

func (y *Yoloer) yolo(draft bool, mergeStrategy string, confirmed bool, preYoloCommand string, trunk string) {
	if !confirmed {
		fmt.Println("Aborted.")
		return
	}
	y.merger.merge(preYoloCommand, false, true, draft, mergeStrategy, "", false)
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "yolo",
		Short: "Merge a change without running checks",
		Long:  "YOLO mode merges changes without running pre-merge checks. Use with caution!",
		Run: func(cmd *cobra.Command, args []string) {
			cfg := config.Load(config.DefaultFileName)

			shell := shell.New()
			vcs := version_control.NewGit(shell)

			fmt.Println("🚀 YOLO mode enabled - using admin on forge to override branch protection!")
			if cfg.PreYoloCommand != "" {
				fmt.Println("🦺 Will run the following command before merging: ", cfg.PreYoloCommand)
			} else {
				fmt.Println("💀 No premergecommand set in config, so no safety net at all!️")
			}
			fmt.Println("🤔 Are you sure you want to merge? (y/n)")
			var confirm string
			fmt.Scanln(&confirm)

			merger := NewMerger(forge.NewGitHub(shell), vcs)

			yoloer := NewYoloer(merger)
			yoloer.yolo(cfg.Draft, cfg.MergeStrategy, confirm == "y", cfg.PreYoloCommand, cfg.Trunk)
		},
	})
}

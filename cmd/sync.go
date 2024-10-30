package cmd

import (
	"fmt"
	"gokid/config"
	"gokid/shell"
	"gokid/version_control"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "sync",
		Short: "Sync current branch with trunk",
		Long:  "Syncs the current branch with the trunk branch by merging trunk into the current branch",
		Run: func(cmd *cobra.Command, args []string) {
			cfg := config.Load(config.DefaultFileName)
			git := version_control.NewGit(shell.New())

			if err := git.SyncTrunk(cfg.Trunk); err != nil {
				fmt.Println("Error syncing with trunk:", err)
				return
			}
			fmt.Printf("Successfully synced with %s\n", cfg.Trunk)
		},
		Aliases: []string{"s"},
	})
}

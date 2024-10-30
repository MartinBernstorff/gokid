package cmd

import (
	"gokid/forge"
	"gokid/shell"

	"github.com/spf13/cobra"
)

func view(forge forge.Forge) {
	forge.ViewPullRequest()
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "view",
		Short: "View the change at the forge",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			view(forge.NewGitHub(shell.New()))
		},
		Aliases: []string{"v"},
	})
}

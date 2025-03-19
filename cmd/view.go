package cmd

import (
	"gokid/forge"
	"gokid/shell"

	"github.com/spf13/cobra"
)

func view(forge forge.Forge) error {
	return forge.ViewPullRequest()
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "view",
		Short: "View the change at the forge",
		Long:  "",
		Run: func(_ *cobra.Command, _ []string) {
			err := view(forge.NewGitHub(shell.New()))
			if err != nil {
				panic(err)
			}
		},
		Aliases: []string{"v"},
	})
}

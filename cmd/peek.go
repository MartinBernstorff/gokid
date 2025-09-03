package cmd

import (
	"gokid/forge"
	"gokid/shell"

	"github.com/spf13/cobra"
)

func peek(forge forge.Forge) error {
	return forge.PeekPullRequest()
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "peek",
		Short: "Peek at the PR in the terminal",
		Long:  "",
		Run: func(_ *cobra.Command, _ []string) {
			err := peek(forge.NewGitHub(shell.New()))
			if err != nil {
				panic(err)
			}
		},
		Aliases: []string{"p"},
	})
}

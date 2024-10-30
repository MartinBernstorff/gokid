package cmd

import (
	"gokid/shell"

	"github.com/spf13/cobra"
)

func markReady() {
	myShell := shell.New()
	myShell.Run("gh pr ready")
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "ready",
		Short: "Mark a change as ready for review",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			markReady()
		},
		Aliases: []string{"r"},
	})
}

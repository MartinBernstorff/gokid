package cmd

import (
	"gokid/forge"

	"github.com/spf13/cobra"
)

func view() {
	forge.View()
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "view",
		Short: "View the change at the forge",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			view()
		},
		Aliases: []string{"v"},
	})
}

/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"gokid/shell"

	"github.com/spf13/cobra"
)

func markDraft() {
	shell.Run("gh pr merge --disable-auto")

}

// draftCmd represents the draft command
var draftCmd = &cobra.Command{
	Use:   "draft",
	Short: "Mark a change as draft",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		markDraft()
	},
	Aliases: []string{"draft"},
}

func init() {
	rootCmd.AddCommand(draftCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// readyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// readyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

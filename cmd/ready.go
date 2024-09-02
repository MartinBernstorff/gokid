/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"gokid/shell"

	"github.com/spf13/cobra"
)

func markReady() {
	// TODO: Implement
	shell.Run("gh pr merge --auto --squash")

}

// readyCmd represents the ready command
var readyCmd = &cobra.Command{
	Use:   "ready",
	Short: "Mark a change as ready for review",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		markReady()
	},
	Aliases: []string{"r"},
}

func init() {
	rootCmd.AddCommand(readyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// readyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// readyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

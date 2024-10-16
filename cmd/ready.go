/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"gokid/shell"

	"github.com/spf13/cobra"
)

func markReady() {
	shell.Run("gh pr ready")
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
}

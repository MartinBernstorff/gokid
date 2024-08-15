/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"gokid/forge"
	"gokid/shell"
	"gokid/vcs"

	"github.com/spf13/cobra"
)

func newChange(inputTitle string) {
	issueTitle := forge.ParseIssueTitle(inputTitle)
	vcs.NewChange(forge.Issue{Title: issueTitle}, "main", true)
	shell.Run("gh", "pr", "create", "--title", issueTitle.Content, "--body", " ")
}

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new change",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		newChange(args[0])
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

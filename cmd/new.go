/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"gokid/forge"
	"gokid/shell"
	"gokid/vcs"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func changeInput() forge.IssueTitle {
	validate := func(input string) error {
		containsPrefix := strings.Contains(input, ":")
		if !containsPrefix {
			return errors.New("invalid change")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Title of change",
		Validate: validate,
	}

	result, _ := prompt.Run()
	return forge.ParseIssueTitle(result)
}

func newChange() {
	issueTitle := changeInput()
	vcs.NewChange(forge.Issue{Title: issueTitle}, "main", true)
	shell.Run("gh pr create --title " + issueTitle.Content + " --body \"\"")
}

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new change",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		newChange()
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

package cmd

import (
	"errors"
	"gokid/config"
	"gokid/forge"
	"gokid/shell"
	"gokid/version_control"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func changeInput() forge.IssueTitle {
	validate := func(input string) error {
		if len(input) < 1 {
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
	cfg := config.Init()

	issueTitle := changeInput()
	version_control.NewChange(forge.Issue{Title: issueTitle}, cfg.Trunk, true)
	shell.Run("gh pr create --title \"" + issueTitle.Prefix + ": " + issueTitle.Content + "\" --body \"\"")
}

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new change",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		newChange()
	},
	Aliases: []string{"n"},
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

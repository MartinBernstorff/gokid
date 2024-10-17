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

	// Handle forge
	cmd := "gh pr create --base " + cfg.Trunk

	if cfg.Draft {
		cmd += " --draft"
	}
	cmd += " --title \"" + issueTitle.Prefix + ": " + issueTitle.Content + "\" --body \"\""

	shell.Run(cmd)
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "new",
		Short: "Create a new change",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			newChange()
		},
		Aliases: []string{"n"},
	})
}

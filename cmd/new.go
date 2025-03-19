package cmd

import (
	"errors"
	"fmt"
	"gokid/commands"
	"gokid/config"
	"gokid/forge"
	"gokid/shell"
	"gokid/version_control"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func changeNamePrompt() string {
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
	return result
}

func newChange(f forge.Forge, cfg *config.GokidConfig, inputTitle string, versionControl version_control.VCS) error {
	parsedTitle := forge.ParseIssueTitle(inputTitle)

	executables := []commands.Command{
		commands.NewFetchOriginCommand(),
		commands.NewCreateBranchCommand(parsedTitle, cfg.Trunk),
		commands.NewEmptyCommitCommand(),
		commands.NewPushCommand(),
	}

	clean, err := versionControl.IsClean()
	if err != nil {
		return err
	}

	if !clean {
		// Add to the stash first
		executables = append([]commands.Command{commands.NewStashCommand()}, executables...)

		// Remember to pop the stash at the end
		executables = append(executables, commands.NewPopStashCommand())
	}

	// Create the PR
	executables = append(executables, commands.NewPullRequestCommand(
		parsedTitle,
		cfg.Trunk,
		cfg.Draft,
	))

	errors := commands.Execute(executables)
	if len(errors) > 0 {
		return errors[0]
	}
	return nil
}

func init() {
	newCmd := &cobra.Command{
		Use:   "new [title]",
		Short: "Create a new change",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			cfg := config.Load(config.DefaultFileName)
			shell := shell.New()

			var title string
			switch len(args) {
			case 0:
				title = changeNamePrompt()
			case 1:
				title = args[0]
			default:
				fmt.Fprintf(os.Stderr, "Error: Too many arguments\n")
				os.Exit(1)
			}

			if err := newChange(forge.NewGitHub(shell), &cfg, title, version_control.NewGit(shell)); err != nil {
				fmt.Fprintf(os.Stderr, "Error creating change: %v\n", err)
				os.Exit(1)
			}
		},
		Aliases: []string{"n"},
	}
	rootCmd.AddCommand(newCmd)
}

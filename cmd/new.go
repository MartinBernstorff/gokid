package cmd

import (
	"errors"
	"fmt"
	"gokid/commands"
	"gokid/config"
	"gokid/forge"
	"gokid/shell"
	"gokid/versioncontrol"
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

func newChange(git versioncontrol.Git, github forge.GitHubForge, cfg *config.GokidConfig, inputTitle string, versionControl versioncontrol.VCS) error {
	parsedTitle := forge.ParseIssueTitle(inputTitle)

	executables := []commands.Command{
		versioncontrol.NewFetchOriginCommand(git),
		versioncontrol.NewCreateBranchCommand(git, parsedTitle, cfg.Trunk),
		versioncontrol.NewEmptyCommitCommand(git),
		versioncontrol.NewPushCommand(git),
	}

	clean, err := versionControl.IsClean()
	if err != nil {
		return err
	}

	if !clean {
		// Add to the stash first
		executables = append([]commands.Command{versioncontrol.NewStashCommand(git)}, executables...)

		// Remember to pop the stash at the end
		executables = append(executables, versioncontrol.NewPopStashCommand(git))
	}

	// Create the PR
	executables = append(executables, forge.NewPullRequestCommand(
		github,
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
		Run: func(_ *cobra.Command, args []string) {
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

			git := versioncontrol.NewGit(shell)
			github := forge.NewGitHub(shell)
			if err := newChange(*git, *github, &cfg, title, versioncontrol.NewGit(shell)); err != nil {
				fmt.Fprintf(os.Stderr, "Error creating change: %v\n", err)
				os.Exit(1)
			}
		},
		Aliases: []string{"n"},
	}
	rootCmd.AddCommand(newCmd)
}

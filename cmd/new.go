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

func changeNamePrompt(label string) string {
	validate := func(input string) error {
		if len(input) < 1 {
			return errors.New("change name cannot be empty")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    label,
		Validate: validate,
	}

	result, _ := prompt.Run()
	return result
}

func NewPrintStatusCommand(status string) commands.Command {
	return commands.Command{
		Assumptions: []commands.NamedCallable{},
		Action: commands.NamedCallable{
			Name: "Print status",
			Callable: func() error {
				fmt.Println(status)
				return nil
			},
		},
		Revert: commands.NamedCallable{},
	}
}

func newChange(git versioncontrol.Git, github forge.GitHubForge, cfg *config.GokidConfig, inputTitle string, description string, versionControl versioncontrol.VCS, commitChanges bool) []error {
	parsedTitle := forge.ParseIssueTitle(inputTitle)
	currentCommit, err := git.Ops.CurrentCommit()
	if err != nil {
		return []error{fmt.Errorf("could not determine current commit: %v", err)}
	}

	executables := []commands.Command{
		versioncontrol.NewCreateBranchCommand(git, parsedTitle, cfg.Trunk),
		versioncontrol.NewEmptyCommitCommand(git),
		NewPrintStatusCommand("Ready to accept commits on: " + parsedTitle.ToBranchName().String()),
		versioncontrol.NewFetchOriginCommand(git, cfg.Trunk),
		versioncontrol.NewRebaseCommand(git, cfg.Trunk, currentCommit),
		versioncontrol.NewPushCommand(git, parsedTitle.ToBranchName()),
	}

	clean, err := versionControl.IsClean()
	if err != nil {
		return []error{fmt.Errorf("could not determine vcs status: %v", err)}
	}

	if !clean && !commitChanges {
		// Add to the stash first
		executables = append([]commands.Command{versioncontrol.NewStashCommand(git)}, executables...)

		// Remember to pop the stash at the end
		executables = append(executables, versioncontrol.NewPopStashCommand(git))
	}

	// Create the PR
	executables = append(executables, forge.NewPullRequestCommand(
		github,
		parsedTitle,
		description,
		cfg.Trunk,
		cfg.Draft,
	))

	errors := commands.Execute(executables)
	if len(errors) > 0 {
		return errors
	}
	return nil
}

func init() {
	var commitChanges bool

	newCmd := &cobra.Command{
		Use:   "new [title]",
		Short: "Create a new change",
		Long:  "",
		Run: func(_ *cobra.Command, args []string) {
			cfg := config.Load(config.DefaultFileName)
			shell := shell.New()

			var title string
			var description string
			switch len(args) {
			case 0:
				title = changeNamePrompt("Title:")
				description = ""
			case 1:
				title = args[0]
				description = ""
			case 2:
				title = args[0]
				description = args[1]
			default:
				fmt.Fprintf(os.Stderr, "too many arguments\n")
				os.Exit(1)
			}

			git := versioncontrol.NewGit(shell)
			github := forge.NewGitHub(shell)
			if err := newChange(*git, *github, &cfg, title, description, versioncontrol.NewGit(shell), commitChanges); err != nil {
				if len(err) != 0 {
					fmt.Fprintf(os.Stderr, "Errors occurred during execution:\n")
					for _, e := range err {
						fmt.Fprintf(os.Stderr, "- %v\n", e)
					}
				}
				// Errors are logged previously
				os.Exit(1)
			}
		},
		Aliases: []string{"n"},
	}

	newCmd.Flags().BoolVar(&commitChanges, "commit", false, "Commit current changes instead of stashing")
	rootCmd.AddCommand(newCmd)
}

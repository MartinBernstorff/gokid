package cmd

import (
	"fmt"
	"gokid/forge"
	"gokid/shell"
	"gokid/versioncontrol"
	"os"

	"github.com/spf13/cobra"
)

func closeChange(github forge.GitHubForge, git versioncontrol.Git, comment string) error {
	branch, err := git.Ops.CurrentBranch()
	if err != nil {
		return fmt.Errorf("error getting current branch: %w", err)
	}

	err = github.CloseChange(comment, branch)
	if err != nil {
		return fmt.Errorf("error closing pull request: %w", err)
	}
	return nil
}

func init() {
	newCmd := &cobra.Command{
		Use:   "close [comment]",
		Short: "Close the current change",
		Long:  "",
		Run: func(_ *cobra.Command, args []string) {
			shell := shell.New()

			var comment string
			switch len(args) {
			case 0:
				comment = ""
			case 1:
				comment = args[0]
			default:
				fmt.Println("Too many arguments")
				os.Exit(1)
			}

			github := forge.NewGitHub(shell)
			git := versioncontrol.NewGit(shell)
			if err := closeChange(*github, *git, comment); err != nil {
				fmt.Fprintf(os.Stderr, "creating change: %v\n", err)
				os.Exit(1)
			}
		},
		Aliases: []string{"c"},
	}
	rootCmd.AddCommand(newCmd)
}

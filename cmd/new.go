package cmd

import (
	"errors"
	"fmt"
	"gokid/config"
	"gokid/forge"
	"gokid/version_control"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func changeInput() string {
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

	if err := versionControl.NewChange(forge.Issue{Title: parsedTitle}, cfg.Trunk, true, cfg.BranchPrefix, cfg.BranchSuffix); err != nil {
		return err
	}

	return f.CreatePullRequest(forge.Issue{Title: parsedTitle}, cfg.Trunk, cfg.Draft)
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "new",
		Short: "Create a new change",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			cfg := config.Load(config.DefaultFileName)
			if err := newChange(forge.NewGitHub(), &cfg, changeInput(), version_control.NewGit()); err != nil {
				fmt.Fprintf(os.Stderr, "Error creating change: %v\n", err)
				os.Exit(1)
			}
		},
		Aliases: []string{"n"},
	})
}

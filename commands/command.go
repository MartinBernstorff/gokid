package commands

import (
	"fmt"
	"gokid/forge"
	"gokid/shell"
	"gokid/versioncontrol"
)

type Command struct {
	description string
	assumptions []func() error
	action      func() error
	revert      func() error
}

func NewFetchOriginCommand() Command {
	return Command{
		description: "Fetch origin",
		assumptions: []func() error{},
		action: func() error {
			// p2: Hacky implementation, should this be a "CreateFetchOriginCommand" which takes the shell as an argument?
			// If we ever need to support more than one forge/vcs, that's definitely the case.
			// p3: Perhaps the commands should be the only thing that's exported, not the methods? If so, the commands need to be in the same package as the methods.
			git := versioncontrol.NewGit(shell.New())
			return git.Ops.Fetch("origin")
		},
		revert: nil,
	}
}

func NewCreateBranchCommand(issueTitle forge.IssueTitle, defaultBranch string) Command {
	// Parse the branch title in the same way as currently
	newBranchName := forge.NewBranchName(issueTitle.Content)

	return Command{
		description: fmt.Sprintf("Create branch %s", newBranchName),
		assumptions: []func() error{
			func() error {
				git := versioncontrol.NewGit(shell.New())
				exists, err := git.Ops.BranchExists(newBranchName.String())
				if err != nil {
					return err
				}

				if exists {
					return fmt.Errorf("branch %s already exists", issueTitle)
				}
				return nil
			},
		},
		action: func() error {
			git := versioncontrol.NewGit(shell.New())
			return git.Ops.BranchFromOrigin(newBranchName.String(), defaultBranch)
		},
		revert: func() error {
			git := versioncontrol.NewGit(shell.New())
			err := git.Ops.SwitchBranch(defaultBranch)
			if err != nil {
				return err
			}
			return git.Ops.DeleteBranch(newBranchName.String())
		},
	}
}

func NewEmptyCommitCommand() Command {
	return Command{
		description: "Create an empty commit",
		assumptions: []func() error{},
		action: func() error {
			git := versioncontrol.NewGit(shell.New())
			return git.Ops.EmptyCommit("Initial commit")
		},
		revert: nil,
	}
}

func NewPushCommand() Command {
	return Command{
		description: "Push to origin",
		assumptions: []func() error{},
		action: func() error {
			git := versioncontrol.NewGit(shell.New())
			return git.Ops.Push()
		},
		revert: nil,
	}
}

func NewStashCommand() Command {
	return Command{
		description: "Stash changes",
		assumptions: []func() error{},
		action: func() error {
			git := versioncontrol.NewGit(shell.New())
			return git.Stash.Save()
		},
		revert: func() error {
			git := versioncontrol.NewGit(shell.New())
			return git.Stash.Pop()
		},
	}
}

func NewPopStashCommand() Command {
	return Command{
		description: "Pop stash",
		assumptions: []func() error{},
		action: func() error {
			git := versioncontrol.NewGit(shell.New())
			return git.Stash.Pop()
		},
		revert: nil,
	}
}

func NewPullRequestCommand(title forge.IssueTitle, trunk string, draft bool) Command {
	f := forge.NewGitHub(shell.New())

	return Command{
		description: fmt.Sprintf("Create pull request %s", title),
		assumptions: []func() error{},
		action: func() error {
			return f.CreatePullRequest(forge.Issue{Title: title}, trunk, draft)
		},
		// p2: Close the PR
		revert: nil,
	}
}

func NewFailCommand() Command {
	return Command{
		description: "Fail",
		assumptions: []func() error{},
		action: func() error {
			return fmt.Errorf("fail")
		},
		revert: nil,
	}
}

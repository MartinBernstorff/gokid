package commands

import (
	"fmt"
	"gokid/forge"
	"gokid/shell"
	"gokid/version_control"
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
			git := version_control.NewGit(shell.New())
			git.Ops.Fetch("origin")
			return nil
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
				git := version_control.NewGit(shell.New())
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
			git := version_control.NewGit(shell.New())
			git.Ops.BranchFromOrigin(newBranchName.String(), defaultBranch)
			return nil
		},
		revert: func() error {
			git := version_control.NewGit(shell.New())
			git.Ops.SwitchBranch(defaultBranch)
			git.Ops.DeleteBranch(newBranchName.String())
			return nil
		},
	}
}

func NewEmptyCommitCommand() Command {
	return Command{
		description: "Create an empty commit",
		assumptions: []func() error{},
		action: func() error {
			git := version_control.NewGit(shell.New())
			git.Ops.EmptyCommit("Initial commit")
			return nil
		},
		revert: nil,
	}
}

func NewPushCommand() Command {
	return Command{
		description: "Push to origin",
		assumptions: []func() error{},
		action: func() error {
			git := version_control.NewGit(shell.New())
			git.Ops.Push()
			return nil
		},
		revert: nil,
	}
}

func NewStashCommand() Command {
	return Command{
		description: "Stash changes",
		assumptions: []func() error{},
		action: func() error {
			git := version_control.NewGit(shell.New())
			git.Stash.Save()
			return nil
		},
		revert: func() error {
			git := version_control.NewGit(shell.New())
			git.Stash.Pop()
			return nil
		},
	}
}

func NewPopStashCommand() Command {
	return Command{
		description: "Pop stash",
		assumptions: []func() error{},
		action: func() error {
			git := version_control.NewGit(shell.New())
			git.Stash.Pop()
			return nil
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
			f.CreatePullRequest(forge.Issue{Title: title}, trunk, draft)
			return nil
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

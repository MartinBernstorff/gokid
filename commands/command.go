package commands

import (
	"fmt"
	"gokid/forge"
	"gokid/shell"
	"gokid/version_control"
)

type Command struct {
	assumptions []func() error
	action      func() error
	revert      func() error
}

func NewFetchOriginCommand() Command {
	return Command{
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
	branchName := version_control.BranchTitle(issueTitle, "", "")

	return Command{
		assumptions: []func() error{
			func() error {
				git := version_control.NewGit(shell.New())
				exists, err := git.Ops.BranchExists(branchName)
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
			git.Ops.BranchFromOrigin(branchName, defaultBranch)
			return nil
		},
		// p2: Delete the branch
		revert: func() error {
			git := version_control.NewGit(shell.New())
			git.Ops.DeleteBranch(branchName)
			return nil
		},
	}
}

func NewEmptyCommitCommand() Command {
	return Command{
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
		assumptions: []func() error{},
		action: func() error {
			f.CreatePullRequest(forge.Issue{Title: title}, trunk, draft)
			return nil
		},
		revert: nil,
	}
}

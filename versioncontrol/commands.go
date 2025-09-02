package versioncontrol

import (
	"fmt"
	"gokid/commands"
	"gokid/forge"
)

func NewFetchOriginCommand(git Git, branch string) commands.Command {
	return commands.Command{
		Assumptions: []commands.NamedCallable{},
		Action: commands.NamedCallable{
			Name: "Fetch origin",
			Callable: func() error {
				return git.Ops.fetch("origin", branch)
			},
		},
		Revert: commands.NamedCallable{},
	}
}

func NewCreateBranchCommand(git Git, issueTitle forge.IssueTitle, defaultBranch string) commands.Command {
	newBranchName := forge.NewBranchName(issueTitle.Content)

	startingBranch, err := git.Ops.CurrentBranch()
	if err != nil {
		panic(err)
	}

	return commands.Command{
		Assumptions: []commands.NamedCallable{
			{
				Name: "Branch does not exist",
				Callable: func() error {
					exists, err := git.Ops.branchExists(newBranchName.String())
					if err != nil {
						return fmt.Errorf("creating branch: %s", err)
					}
					if exists {
						return fmt.Errorf("branch %s already exists", issueTitle)
					}
					return nil
				},
			},
		},
		Action: commands.NamedCallable{
			Name: fmt.Sprintf("Branch out from trunk with '%s'", newBranchName),
			Callable: func() error {
				return git.Ops.branchFromOrigin(newBranchName.String(), defaultBranch)
			},
		},
		Revert: commands.NamedCallable{
			Name: "Delete branch",
			Callable: func() error {
				err := git.Ops.switchBranch(startingBranch)
				if err != nil {
					return err
				}
				return git.Ops.deleteBranch(newBranchName.String())
			},
		},
	}
}

func NewEmptyCommitCommand(git Git) commands.Command {
	return commands.Command{
		Assumptions: []commands.NamedCallable{},
		Action: commands.NamedCallable{
			Name: "Create an empty commit",
			Callable: func() error {
				return git.Ops.emptyCommit("Initial commit")
			},
		},
		Revert: commands.NamedCallable{},
	}
}

func NewPushCommand(git Git, branchName forge.BranchName) commands.Command {
	return commands.Command{
		Assumptions: []commands.NamedCallable{},
		Action: commands.NamedCallable{
			Name: "Push",
			Callable: func() error {
				return git.Ops.push(branchName)
			},
		},
		Revert: commands.NamedCallable{},
	}
}

func NewStashCommand(git Git) commands.Command {
	return commands.Command{
		Assumptions: []commands.NamedCallable{},
		Action: commands.NamedCallable{
			Name: "Stash changes",
			Callable: func() error {
				return git.Stash.Save()
			},
		},
		Revert: commands.NamedCallable{
			Name: "Pop stash",
			Callable: func() error {
				return git.Stash.Pop()
			},
		},
	}
}

func NewPopStashCommand(git Git) commands.Command {
	return commands.Command{
		Assumptions: []commands.NamedCallable{},
		Action: commands.NamedCallable{
			Name: "Pop stash",
			Callable: func() error {
				return git.Stash.Pop()
			},
		},
		Revert: commands.NamedCallable{
			Name: "Stash changes",
			Callable: func() error {
				return git.Stash.Save()
			},
		},
	}
}

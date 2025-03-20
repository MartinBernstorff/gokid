package versioncontrol

import (
	"fmt"
	"gokid/commands"
	"gokid/forge"
)

func NewFetchOriginCommand(git Git) commands.Command {
	return commands.Command{
		Assumptions: []commands.NamedCallable{},
		Action: commands.NamedCallable{
			Name: "Fetch origin",
			Callable: func() error {
				return git.ops.fetch("origin")
			},
		},
		Revert: commands.NamedCallable{},
	}
}

func NewCreateBranchCommand(git Git, issueTitle forge.IssueTitle, defaultBranch string) commands.Command {
	newBranchName := forge.NewBranchName(issueTitle.Content)

	startingBranch, err := git.ops.currentBranch()
	if err != nil {
		panic(err)
	}

	return commands.Command{
		Assumptions: []commands.NamedCallable{
			{
				Name: fmt.Sprintf("Branch does not exist", newBranchName),
				Callable: func() error {
					exists, err := git.ops.branchExists(newBranchName.String())
					if err != nil {
						return err
					}
					if exists {
						return fmt.Errorf("branch %s already exists", issueTitle)
					}
					return nil
				},
			},
		},
		Action: commands.NamedCallable{
			Name: fmt.Sprintf("Create branch '%s'", newBranchName),
			Callable: func() error {
				return git.ops.branchFromOrigin(newBranchName.String(), defaultBranch)
			},
		},
		Revert: commands.NamedCallable{
			Name: "Delete branch",
			Callable: func() error {
				err := git.ops.switchBranch(startingBranch)
				if err != nil {
					return err
				}
				return git.ops.deleteBranch(newBranchName.String())
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
				return git.ops.emptyCommit("Initial commit")
			},
		},
		Revert: commands.NamedCallable{},
	}
}

func NewPushCommand(git Git) commands.Command {
	return commands.Command{
		Assumptions: []commands.NamedCallable{},
		Action: commands.NamedCallable{
			Name: "Push",
			Callable: func() error {
				return git.ops.push()
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

package commands

import (
	"fmt"
	"gokid/forge"
	"gokid/versioncontrol"
)

type Command struct {
	assumptions []NamedCallable
	action      NamedCallable
	revert      NamedCallable
}

type NamedCallable struct {
	name     string
	callable func() error
}

// p3: Perhaps the commands should be the only thing that's exported, not the methods? If so, the commands need to be in the same package as the methods.

func NewFetchOriginCommand(git versioncontrol.Git) Command {
	return Command{
		assumptions: []NamedCallable{},
		action: NamedCallable{
			name: "fetch origin",
			callable: func() error {
				return git.Ops.Fetch("origin")
			},
		},
		revert: NamedCallable{},
	}
}

func NewCreateBranchCommand(git versioncontrol.Git, issueTitle forge.IssueTitle, defaultBranch string) Command {
	newBranchName := forge.NewBranchName(issueTitle.Content)

	startingBranch, err := git.Ops.CurrentBranch()
	if err != nil {
		panic(err)
	}

	return Command{
		assumptions: []NamedCallable{
			{
				name: fmt.Sprintf("%s does not exist", newBranchName),
				callable: func() error {
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
		},
		action: NamedCallable{
			name: fmt.Sprintf("Create branch %s", newBranchName),
			callable: func() error {
				return git.Ops.BranchFromOrigin(newBranchName.String(), defaultBranch)
			},
		},
		revert: NamedCallable{
			name: "Delete branch",
			callable: func() error {
				err := git.Ops.SwitchBranch(startingBranch)
				if err != nil {
					return err
				}
				return git.Ops.DeleteBranch(newBranchName.String())
			},
		},
	}
}

func NewEmptyCommitCommand(git versioncontrol.Git) Command {
	return Command{
		assumptions: []NamedCallable{},
		action: NamedCallable{
			name: "Create an empty commit",
			callable: func() error {
				return git.Ops.EmptyCommit("Initial commit")
			},
		},
		revert: NamedCallable{},
	}
}

func NewPushCommand(git versioncontrol.Git) Command {
	return Command{
		assumptions: []NamedCallable{},
		action: NamedCallable{
			name: "Push",
			callable: func() error {
				return git.Ops.Push()
			},
		},
		revert: NamedCallable{},
	}
}

func NewStashCommand(git versioncontrol.Git) Command {
	return Command{
		assumptions: []NamedCallable{},
		action: NamedCallable{
			name: "Stash changes",
			callable: func() error {
				return git.Stash.Save()
			},
		},
		revert: NamedCallable{
			name: "Pop stash",
			callable: func() error {
				return git.Stash.Pop()
			},
		},
	}
}

func NewPopStashCommand(git versioncontrol.Git) Command {
	return Command{
		assumptions: []NamedCallable{},
		action: NamedCallable{
			name: "Pop stash",
			callable: func() error {
				return git.Stash.Pop()
			},
		},
		revert: NamedCallable{
			name: "Stash changes",
			callable: func() error {
				return git.Stash.Save()
			},
		},
	}
}

func NewPullRequestCommand(f forge.GitHubForge, title forge.IssueTitle, trunk string, draft bool) Command {
	return Command{
		assumptions: []NamedCallable{},
		action: NamedCallable{
			name: fmt.Sprintf("Create pull request '%s'", title),
			callable: func() error {
				return f.CreatePullRequest(forge.Issue{Title: title}, trunk, draft)
			},
		},
		// p5: Close the pull request
		revert: NamedCallable{},
	}
}

func NewFailCommand() Command {
	return Command{
		assumptions: []NamedCallable{},
		action: NamedCallable{
			name: "Fail",
			callable: func() error {
				return fmt.Errorf("fail")
			},
		},
		revert: NamedCallable{},
	}
}

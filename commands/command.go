package commands

import (
	"fmt"
	"gokid/forge"
	"gokid/versioncontrol"
)

type Command struct {
	assumptions []LabelledCallable
	action      LabelledCallable
	revert      LabelledCallable
}

type LabelledCallable struct {
	name     string
	callable func() error
}

// p3: Perhaps the commands should be the only thing that's exported, not the methods? If so, the commands need to be in the same package as the methods.

func NewFetchOriginCommand(git versioncontrol.Git) Command {
	return Command{
		assumptions: []LabelledCallable{},
		action: LabelledCallable{
			name: "fetch origin",
			callable: func() error {
				return git.Ops.Fetch("origin")
			},
		},
		revert: LabelledCallable{},
	}
}

func NewCreateBranchCommand(git versioncontrol.Git, issueTitle forge.IssueTitle, defaultBranch string) Command {
	newBranchName := forge.NewBranchName(issueTitle.Content)

	startingBranch, err := git.Ops.CurrentBranch()
	if err != nil {
		panic(err)
	}

	return Command{
		assumptions: []LabelledCallable{
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
		action: LabelledCallable{
			name: fmt.Sprintf("Create branch %s", newBranchName),
			callable: func() error {
				return git.Ops.BranchFromOrigin(newBranchName.String(), defaultBranch)
			},
		},
		revert: LabelledCallable{
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
		assumptions: []LabelledCallable{},
		action: LabelledCallable{
			name: "Create an empty commit",
			callable: func() error {
				return git.Ops.EmptyCommit("Initial commit")
			},
		},
		revert: LabelledCallable{},
	}
}

func NewPushCommand(git versioncontrol.Git) Command {
	return Command{
		assumptions: []LabelledCallable{},
		action: LabelledCallable{
			name: "Push",
			callable: func() error {
				return git.Ops.Push()
			},
		},
		revert: LabelledCallable{},
	}
}

func NewStashCommand(git versioncontrol.Git) Command {
	return Command{
		assumptions: []LabelledCallable{},
		action: LabelledCallable{
			name: "Stash changes",
			callable: func() error {
				return git.Stash.Save()
			},
		},
		revert: LabelledCallable{
			name: "Pop stash",
			callable: func() error {
				return git.Stash.Pop()
			},
		},
	}
}

func NewPopStashCommand(git versioncontrol.Git) Command {
	return Command{
		assumptions: []LabelledCallable{},
		action: LabelledCallable{
			name: "Pop stash",
			callable: func() error {
				return git.Stash.Pop()
			},
		},
		revert: LabelledCallable{
			name: "Stash changes",
			callable: func() error {
				return git.Stash.Save()
			},
		},
	}
}

func NewPullRequestCommand(f forge.GitHubForge, title forge.IssueTitle, trunk string, draft bool) Command {
	return Command{
		assumptions: []LabelledCallable{},
		action: LabelledCallable{
			name: fmt.Sprintf("Create pull request %s", title),
			callable: func() error {
				return f.CreatePullRequest(forge.Issue{Title: title}, trunk, draft)
			},
		},
		// p3: Close the pull request
		revert: LabelledCallable{},
	}
}

func NewFailCommand() Command {
	return Command{
		assumptions: []LabelledCallable{},
		action: LabelledCallable{
			name: "Fail",
			callable: func() error {
				return fmt.Errorf("fail")
			},
		},
		revert: LabelledCallable{},
	}
}

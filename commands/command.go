package commands

import (
	"fmt"
	"gokid/forge"
	"gokid/shell"
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

// p2: Hacky implementation, should this be a "CreateFetchOriginCommand" which takes the shell as an argument?
// If we ever need to support more than one forge/vcs, that's definitely the case.

// p3: Perhaps the commands should be the only thing that's exported, not the methods? If so, the commands need to be in the same package as the methods.

func NewFetchOriginCommand() Command {
	return Command{
		assumptions: []LabelledCallable{},
		action: LabelledCallable{
			name: "fetch origin",
			callable: func() error {
				git := versioncontrol.NewGit(shell.New())
				return git.Ops.Fetch("origin")
			},
		},
		revert: LabelledCallable{},
	}
}

func NewCreateBranchCommand(issueTitle forge.IssueTitle, defaultBranch string) Command {
	newBranchName := forge.NewBranchName(issueTitle.Content)

	return Command{
		assumptions: []LabelledCallable{
			{
				name: fmt.Sprintf("%s does not exist", newBranchName),
				callable: func() error {
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
		},
		action: LabelledCallable{
			name: fmt.Sprintf("Create branch %s", newBranchName),
			callable: func() error {
				git := versioncontrol.NewGit(shell.New())
				return git.Ops.BranchFromOrigin(newBranchName.String(), defaultBranch)
			},
		},
		revert: LabelledCallable{
			name: "Delete branch",
			callable: func() error {
				git := versioncontrol.NewGit(shell.New())

				// p2: Switch to the branch we were on when we created the command
				err := git.Ops.SwitchBranch(defaultBranch)
				if err != nil {
					return err
				}
				return git.Ops.DeleteBranch(newBranchName.String())
			},
		},
	}
}

func NewEmptyCommitCommand() Command {
	return Command{
		assumptions: []LabelledCallable{},
		action: LabelledCallable{
			name: "Create an empty commit",
			callable: func() error {
				git := versioncontrol.NewGit(shell.New())
				return git.Ops.EmptyCommit("Initial commit")
			},
		},
		revert: LabelledCallable{},
	}
}

func NewPushCommand() Command {
	return Command{
		assumptions: []LabelledCallable{},
		action: LabelledCallable{
			name: "Push",
			callable: func() error {
				git := versioncontrol.NewGit(shell.New())
				return git.Ops.Push()
			},
		},
		revert: LabelledCallable{},
	}
}

func NewStashCommand() Command {
	return Command{
		assumptions: []LabelledCallable{},
		action: LabelledCallable{
			name: "Stash changes",
			callable: func() error {
				git := versioncontrol.NewGit(shell.New())
				return git.Stash.Save()
			},
		},
		revert: LabelledCallable{
			name: "Pop stash",
			callable: func() error {
				git := versioncontrol.NewGit(shell.New())
				return git.Stash.Pop()
			},
		},
	}
}

func NewPopStashCommand() Command {
	return Command{
		assumptions: []LabelledCallable{},
		action: LabelledCallable{
			name: "Pop stash",
			callable: func() error {
				git := versioncontrol.NewGit(shell.New())
				return git.Stash.Pop()
			},
		},
		revert: LabelledCallable{
			name: "Stash changes",
			callable: func() error {
				git := versioncontrol.NewGit(shell.New())
				return git.Stash.Save()
			},
		},
	}
}

func NewPullRequestCommand(title forge.IssueTitle, trunk string, draft bool) Command {
	f := forge.NewGitHub(shell.New())

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

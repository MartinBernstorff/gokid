package forge

import (
	"fmt"
	"gokid/commands"
)

func NewPullRequestCommand(f GitHubForge, title IssueTitle, description string, trunk string, draft bool) commands.Command {
	return commands.Command{
		Assumptions: []commands.NamedCallable{},
		Action: commands.NamedCallable{
			Name: fmt.Sprintf("Create pull request '%s'", title),
			Callable: func() error {
				return f.CreatePullRequest(Issue{Title: title, Body: description}, trunk, draft)
			},
		},
		// p5: Close the pull request
		// Very low priority, because pull request creation is the last step in the process,
		// so if it fails, there's nothing to revert.
		Revert: commands.NamedCallable{},
	}
}

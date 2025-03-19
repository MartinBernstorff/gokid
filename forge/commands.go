package forge

import (
	"fmt"
	"gokid/commands"
)

func NewPullRequestCommand(f GitHubForge, title IssueTitle, trunk string, draft bool) commands.Command {
	return commands.Command{
		Assumptions: []commands.NamedCallable{},
		Action: commands.NamedCallable{
			Name: fmt.Sprintf("Create pull request '%s'", title),
			Callable: func() error {
				return f.CreatePullRequest(Issue{Title: title}, trunk, draft)
			},
		},
		// p5: Close the pull request
		Revert: commands.NamedCallable{},
	}
}

package forge

import (
	"fmt"
	"gokid/shell"
)

type GitHub struct{}

func NewGitHub() *GitHub {
	return &GitHub{}
}

func (g *GitHub) CreatePullRequest(issue Issue, base string, draft bool) error {
	cmd := fmt.Sprintf("gh pr create --base %s", base)

	if draft {
		cmd += " --draft"
	}

	cmd += fmt.Sprintf(" --title \"%s\" --body \"\"", issue.Title.String())

	return shell.Run(cmd)
}

func (g *GitHub) ViewPullRequest() error {
	return shell.Run("gh pr view -w")
}

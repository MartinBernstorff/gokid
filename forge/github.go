package forge

import (
	"fmt"
	"gokid/shell"
)

type GitHubForge struct{}

func NewGitHub() *GitHubForge {
	return &GitHubForge{}
}

func (g *GitHubForge) CreatePullRequest(issue Issue, base string, draft bool) error {
	cmd := fmt.Sprintf("gh pr create --base %s", base)

	if draft {
		cmd += " --draft"
	}

	cmd += fmt.Sprintf(" --title \"%s\" --body \"\"", issue.Title.String())

	return shell.Run(cmd)
}

func (g *GitHubForge) ViewPullRequest() error {
	return shell.Run("gh pr view -w")
}

func (g *GitHubForge) MarkPullRequestReady() error {
	return shell.Run("gh pr ready")
}

func (g *GitHubForge) MergePullRequest(strategy string, autoMerge bool) error {
	cmd := fmt.Sprintf("gh pr merge --%s", strategy)
	if autoMerge {
		cmd += " --auto"
	}
	return shell.Run(cmd)
}

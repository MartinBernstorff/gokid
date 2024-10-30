package forge

import (
	"fmt"
	"gokid/shell"
)

type GitHubForge struct {
	shell shell.Shell
}

func NewGitHub() *GitHubForge {
	return &GitHubForge{
		shell: shell.New(),
	}
}

func (g *GitHubForge) CreatePullRequest(issue Issue, base string, draft bool) error {
	cmd := fmt.Sprintf("gh pr create --base %s", base)

	if draft {
		cmd += " --draft"
	}

	cmd += fmt.Sprintf(" --title \"%s\" --body \"\"", issue.Title.String())

	return g.shell.Run(cmd)
}

func (g *GitHubForge) ViewPullRequest() error {
	return g.shell.Run("gh pr view -w")
}

func (g *GitHubForge) MarkPullRequestReady() error {
	return g.shell.Run("gh pr ready")
}

func (g *GitHubForge) MergePullRequest(strategy string, autoMerge bool, forceMerge bool) error {
	cmd := fmt.Sprintf("gh pr merge --%s", strategy)
	if autoMerge {
		cmd += " --auto"
	}
	if forceMerge {
		cmd += " --admin"
	}

	return g.shell.Run(cmd)
}

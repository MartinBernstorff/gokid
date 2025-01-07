package forge

import (
	"fmt"
	"gokid/shell"
)

type GitHubForge struct {
	shell shell.Shell
}

func NewGitHub(s shell.Shell) *GitHubForge {
	return &GitHubForge{
		shell: s,
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
	return g.shell.Run("gh pr view --web || gh repo view --web")
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

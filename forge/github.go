package forge

import (
	"encoding/json"
	"fmt"
	"strings"

	"gokid/shell"
)

type GitHubForge struct {
	shell shell.Shell
}

func (g *GitHubForge) CloseChange(comment string, branch string) error {
	cmd := fmt.Sprintf("gh pr close \"%s\"", branch)
	if comment != "" {
		cmd += fmt.Sprintf(" --comment \"%s\"", comment)
	}
	_, err := g.shell.Run(cmd)
	if err != nil {
		return fmt.Errorf("github CLI errored: %w", err)
	}
	return nil
}

func NewGitHub(s shell.Shell) *GitHubForge {
	// Check if GitHub CLI is installed
	_, err := s.RunQuietly("gh --version")
	if err != nil {
		panic("GitHub CLI (gh) is not installed or not available in PATH")
	}
	return &GitHubForge{
		shell: s,
	}
}

func (g *GitHubForge) CreatePullRequest(issue Issue, base string, draft bool) error {
	cmd := fmt.Sprintf("gh pr create --base %s", base)

	if draft {
		cmd += " --draft"
	}

	cmd += fmt.Sprintf(" --title \"%s\" --body \"%s\"", issue.Title.String(), issue.Body)

	output, err := g.shell.RunQuietly(cmd)
	if err != nil {
		return fmt.Errorf("error creating pull request: %s", output)
	}

	fmt.Println("Created pull-request: ", output)

	return nil
}

func (g *GitHubForge) ViewPullRequest() error {
	_, err := g.shell.Run("gh pr view --web || gh repo view --web")
	if err != nil {
		return fmt.Errorf("error viewing pull request: %s", err)
	}
	return nil
}

func (g *GitHubForge) MarkPullRequestReady() error {
	_, err := g.shell.Run("gh pr ready")
	if err != nil {
		return fmt.Errorf("error marking pull request as ready: %s", err)
	}
	return nil
}

func (g *GitHubForge) MergePullRequest(strategy string, autoMerge bool, forceMerge bool) error {
	cmd := fmt.Sprintf("gh pr merge --%s", strategy)
	if autoMerge {
		cmd += " --auto"
	}
	if forceMerge {
		cmd += " --admin"
	}

	_, err := g.shell.RunQuietly(cmd)

	if err != nil {
		return fmt.Errorf("error merging pull request: %s", err)
	}
	return nil
}

// ListPullRequests returns open PRs for the current repo including CI status
func (g *GitHubForge) ListPullRequests() ([]PullRequest, error) {
	// Use gh to list PRs with JSON output. The fields chosen include CI status and URL.
	// Note: statusCheckRollup is available via the GraphQL backing API used by gh.
	fields := []string{"number", "title", "state", "url", "statusCheckRollup"}
	cmd := fmt.Sprintf("gh pr list --state open --json %s", strings.Join(fields, ","))

	out, err := g.shell.RunQuietly(cmd)
	if err != nil {
		return nil, fmt.Errorf("error listing pull requests: %s", out)
	}

	var prs []PullRequest
	if err := json.Unmarshal([]byte(out), &prs); err != nil {
		return nil, fmt.Errorf("failed to parse gh output: %w", err)
	}
	return prs, nil
}

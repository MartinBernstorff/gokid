package forge

import (
	"gokid/shell"
	"regexp"
	"strings"
)

type IssueViewer struct {
	shell shell.Shell
}

func NewIssueViewer(s shell.Shell) *IssueViewer {
	return &IssueViewer{
		shell: s,
	}
}

func (iv *IssueViewer) View() {
	iv.shell.Run("gh pr view -w")
}

type IssueTitle struct {
	Prefix  string
	Content string
}

func ParseIssueTitle(issueTitle string) IssueTitle {
	// Get all string between start and first ":" or "("
	prefixMatch := regexp.MustCompile(`^(.*?)[:]`).FindStringSubmatch(issueTitle)

	if len(prefixMatch) < 2 {
		// No prefix found, return without prefix
		return IssueTitle{Prefix: "", Content: issueTitle}
	}

	prefix := strings.TrimSpace(prefixMatch[1])

	// Get the description after ":"
	descriptionRegex := regexp.MustCompile(`: (.*)$`)
	descriptionMatch := descriptionRegex.FindStringSubmatch(issueTitle)

	if len(descriptionMatch) < 2 {
		// No description found, return with only prefix
		return IssueTitle{Prefix: prefix, Content: ""}
	}

	description := strings.TrimSpace(descriptionMatch[1])

	return IssueTitle{Prefix: prefix, Content: description}
}

func (i IssueTitle) String() string {
	if i.Prefix == "" || i.Content == "" {
		return i.Prefix + i.Content
	}
	return i.Prefix + ": " + i.Content
}

func (i IssueTitle) ToBranchName() (BranchName, error) {
	return NewBranchName(i.Content)
}

func ValidateBranch(branch string) error {
	s := shell.New()
	_, err := s.Run("git check-ref-format --branch " + branch)
	return err
}

type Issue struct {
	Title IssueTitle
	Body  string
}

type RemoteIssue struct {
	Title    IssueTitle
	EntityID string
}

type LocalIssue struct {
	Title IssueTitle
}

type BranchName string

func NewBranchName(name string) (BranchName, error) {
	replacer := strings.NewReplacer(" ", "-", ":", "-", "/", "-", "..", "-", "(", "", ")", "")
	err := ValidateBranch(replacer.Replace(name))
	if err != nil {
		return "", err
	}
	return BranchName(name), nil
}

func (b BranchName) String() string {
	return string(b)
}

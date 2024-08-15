package forge

import (
	"regexp"
	"strings"
)

func ParseIssueTitle(issueTitle string) IssueTitle {
	// Get all string between start and first ":" or "("
	prefixMatch := regexp.MustCompile(`^(.*?)[\(:]`).FindStringSubmatch(issueTitle)

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

type IssueTitle struct {
	Prefix  string
	Content string
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

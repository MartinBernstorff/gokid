package forge

// Forge represents a code hosting platform interface
type Forge interface {
	CreatePullRequest(issue Issue, base string, draft bool) error
	ViewPullRequest() error
	PeekPullRequest() error
	MarkPullRequestReady() error
	MergePullRequest(strategy string, autoMerge bool, forceMerge bool) error
}

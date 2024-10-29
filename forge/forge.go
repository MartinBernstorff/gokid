package forge

// Forge represents a code hosting platform interface
type Forge interface {
	CreatePullRequest(issue Issue, base string, draft bool) error
	ViewPullRequest() error
}

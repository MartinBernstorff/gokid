package forge

// PullRequest models a PR with key fields used across implementations
type PullRequest struct {
	Number int    `json:"number,omitempty"`
	Title  string `json:"title,omitempty"`
	State  string `json:"state,omitempty"`
	URL    string `json:"url,omitempty"`
	Base   string `json:"-"`
	Draft  bool   `json:"isDraft,omitempty"`

	// CI status information from GitHub GraphQL via gh --json statusCheckRollup
	// gh returns an array of either CheckRun or StatusContext items
	StatusCheckRollup []struct {
		TypeName     string `json:"__typename"`
		Name         string `json:"name"`
		Status       string `json:"status"`     // e.g., IN_PROGRESS, COMPLETED
		Conclusion   string `json:"conclusion"` // e.g., SUCCESS, FAILURE, SKIPPED
		WorkflowName string `json:"workflowName"`
		DetailsURL   string `json:"detailsUrl"`
	} `json:"statusCheckRollup,omitempty"`
}

package forge

type FakeForge struct {
	LastCreatedPR     PullRequest
	PRs               []PullRequest
	LastMergeStrategy string
	LastAutoMerge     bool
	LastForceMerge    bool
	WasMarkedReady    bool
}

func NewFakeForge() *FakeForge {
	return &FakeForge{
		PRs: make([]PullRequest, 0),
	}
}

func (f *FakeForge) CreatePullRequest(issue Issue, _ string, draft bool) error {
	pr := PullRequest{}
	pr.Title = issue.Title.String()
	pr.Draft = draft

	f.LastCreatedPR = pr
	f.PRs = append(f.PRs, pr)
	return nil
}

func (f *FakeForge) ViewPullRequest() error {
	return nil
}

func (f *FakeForge) MarkPullRequestReady() error {
	f.WasMarkedReady = true
	return nil
}

func (f *FakeForge) MergePullRequest(strategy string, autoMerge bool, forceMerge bool) error {
	f.LastMergeStrategy = strategy
	f.LastAutoMerge = autoMerge
	f.LastForceMerge = forceMerge
	return nil
}

func (f *FakeForge) ListPullRequests() ([]PullRequest, error) {
	return f.PRs, nil
}

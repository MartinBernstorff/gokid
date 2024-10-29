package forge

type FakeForge struct {
	LastCreatedPR     PullRequest
	PRs               []PullRequest
	LastMergeStrategy string
	LastAutoMerge     bool
	WasMarkedReady    bool
}

type PullRequest struct {
	Title string
	Base  string
	Draft bool
}

func NewFakeForge() *FakeForge {
	return &FakeForge{
		PRs: make([]PullRequest, 0),
	}
}

func (f *FakeForge) CreatePullRequest(issue Issue, base string, draft bool) error {
	pr := PullRequest{
		Title: issue.Title.String(),
		Base:  base,
		Draft: draft,
	}

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

func (f *FakeForge) MergePullRequest(strategy string, autoMerge bool) error {
	f.LastMergeStrategy = strategy
	f.LastAutoMerge = autoMerge
	return nil
}

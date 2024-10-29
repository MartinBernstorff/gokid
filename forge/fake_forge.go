package forge

type FakeForge struct {
	LastCreatedPR PullRequest
	PRs           []PullRequest
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

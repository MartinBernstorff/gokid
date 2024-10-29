package version_control

// StubStash maintains a simple stash counter
type StubStash struct {
	stashCount int
}

func NewStubStash() *StubStash {
	return &StubStash{}
}

func (s *StubStash) Save() {
	s.stashCount++
}

func (s *StubStash) Pop() {
	if s.stashCount > 0 {
		s.stashCount--
	}
}

// Commit represents a git commit with minimal information
type Commit struct {
	Title string
	Empty bool
}

// GitStub simulates minimal git repository state
type GitStub struct {
	BaseGit

	// Repository state
	currentBranch string
	originBranch  string
	isDirty       bool
	commits       []Commit
	lastPush      Commit
}

func NewGitStub() *GitStub {
	g := &GitStub{
		originBranch: "main", // default origin branch
		commits:      make([]Commit, 0),
	}
	g.ops = g
	g.stash = NewStubStash()
	return g
}

// Helper methods to inspect repository state
func (g *GitStub) CurrentBranch() string {
	return g.currentBranch
}

func (g *GitStub) OriginBranch() string {
	return g.originBranch
}

func (g *GitStub) IsDirty() bool {
	return g.isDirty
}

func (g *GitStub) StashCount() int {
	return g.stash.(*StubStash).stashCount
}

func (g *GitStub) Commits() []Commit {
	return g.commits
}

// Implementation of gitOperations interface
func (g *GitStub) SetDirty(isDirty bool) {
	g.isDirty = isDirty
}

func (g *GitStub) isClean() bool {
	return !g.isDirty
}

func (g *GitStub) fetch(remote string) {
}

func (g *GitStub) branchFromOrigin(branchName string, defaultBranch string) {
	g.currentBranch = branchName
	g.originBranch = defaultBranch
}

func (g *GitStub) emptyCommit(message string) {
	g.commits = append(g.commits, Commit{
		Title: message,
		Empty: true,
	})
}

func (g *GitStub) push() {
	g.lastPush = g.commits[len(g.commits)-1]
}

func (g *GitStub) AddCommit(title string, empty bool) {
	g.commits = append(g.commits, Commit{
		Title: title,
		Empty: empty,
	})
}

func (g *GitStub) remoteIsUpdated() bool {
	return g.lastPush.Title == g.commits[len(g.commits)-1].Title
}

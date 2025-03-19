package version_control

// FakeStash maintains a simple stash counter and manages dirty state
type FakeStash struct {
	stashCount int
	git        *FakeGit // Reference to parent git to manage dirty state
}

func NewFakeStash(git *FakeGit) *FakeStash {
	return &FakeStash{
		git: git,
	}
}

func (s *FakeStash) Save() {
	s.stashCount++
	s.git.isDirty = false // Stashing makes working directory clean
}

func (s *FakeStash) Pop() {
	if s.stashCount > 0 {
		s.stashCount--
		s.git.isDirty = true // Popping makes working directory dirty again
	}
}

// Commit represents a git commit with minimal information
type Commit struct {
	Title string
	Empty bool
}

// FakeGit simulates minimal git repository state
type FakeGit struct {
	BaseGit

	// Repository state
	currentBranch    string
	originBranch     string
	isDirty          bool
	commits          []Commit
	lastPush         Commit
	isFetched        bool
	TrunkSynced      bool
	DiffSummaryCalls int
}

func NewFakeGit() *FakeGit {
	g := &FakeGit{
		originBranch:     "main", // default origin branch
		commits:          make([]Commit, 0),
		isFetched:        false,
		TrunkSynced:      false,
		DiffSummaryCalls: 0,
	}
	g.Ops = g
	g.stash = NewFakeStash(g) // Pass git reference to stash
	return g
}

// Helper methods to inspect repository state
func (g *FakeGit) CurrentBranch() string {
	return g.currentBranch
}

func (g *FakeGit) OriginBranch() string {
	return g.originBranch
}

func (g *FakeGit) IsDirty() bool {
	return g.isDirty
}

func (g *FakeGit) StashCount() int {
	return g.stash.(*FakeStash).stashCount
}

func (g *FakeGit) Commits() []Commit {
	return g.commits
}

// Implementation of gitOperations interface
func (g *FakeGit) SetDirty(isDirty bool) {
	g.isDirty = isDirty
}

func (g *FakeGit) IsClean() bool {
	return !g.isDirty
}

func (g *FakeGit) Fetch(remote string) {
	g.isFetched = true
}

func (g *FakeGit) BranchFromOrigin(branchName string, origin string) {
	g.currentBranch = branchName
	g.originBranch = origin
}

func (g *FakeGit) EmptyCommit(message string) {
	g.commits = append(g.commits, Commit{
		Title: message,
		Empty: true,
	})
}

func (g *FakeGit) Push() {
	g.lastPush = g.commits[len(g.commits)-1]
}

func (g *FakeGit) AddCommit(title string, empty bool) {
	g.commits = append(g.commits, Commit{
		Title: title,
		Empty: empty,
	})
}

func (g *FakeGit) remoteIsUpdated() bool {
	return g.lastPush.Title == g.commits[len(g.commits)-1].Title
}

func (g *FakeGit) SyncTrunk(defaultBranch string) error {
	g.TrunkSynced = true
	return nil
}

func (g *FakeGit) ShowDiffSummary(branch string) error {
	g.DiffSummaryCalls += 1
	return nil
}

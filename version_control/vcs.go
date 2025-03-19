package version_control

// VCS defines the interface for version control operations
type VCS interface {
	SyncTrunk(defaultBranch string) error
	ShowDiffSummary(branch string) error
	IsClean() (bool, error)
}

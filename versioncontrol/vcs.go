package versioncontrol

// VCS defines the interface for version control operations
type VCS interface {
	IsClean() (bool, error)
}

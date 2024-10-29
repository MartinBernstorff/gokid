package version_control

import "gokid/forge"

type VCS interface {
	IsClean() bool
	StashChanges()
	PopStash()
	FetchOrigin()
	CheckoutNewBranch(branchName string, baseBranch string)
	CreateEmptyCommit(message string)
	Push()
	NewChange(issue forge.Issue, defaultBranch string, migrateChanges bool, branchPrefix string, branchSuffix string) error
}

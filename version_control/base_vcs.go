package version_control

import (
	"gokid/forge"
	"strings"
)

type BaseVCS struct {
	impl VCS
}

func NewBaseVCS(implementation VCS) *BaseVCS {
	return &BaseVCS{
		impl: implementation,
	}
}

func (b *BaseVCS) branchTitle(issue forge.Issue, prefix string, suffix string) string {
	title := prefix + issue.Title.Content + suffix
	return strings.ReplaceAll(title, " ", "-")
}

func (b *BaseVCS) NewChange(issue forge.Issue, defaultBranch string, migrateChanges bool, branchPrefix string, branchSuffix string) error {
	needsMigration := migrateChanges && !b.impl.IsClean()

	if needsMigration {
		b.impl.StashChanges()
	}

	branchTitle := b.branchTitle(issue, branchPrefix, branchSuffix)
	b.impl.FetchOrigin()
	b.impl.CheckoutNewBranch(branchTitle, defaultBranch)

	if needsMigration {
		b.impl.PopStash()
	}

	b.impl.CreateEmptyCommit(branchTitle)
	b.impl.Push()
	return nil
}

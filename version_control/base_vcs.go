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

func (b *BaseVCS) formatTitle(issue forge.Issue, prefix string, suffix string) string {
	return prefix + issue.Title.Content + suffix
}

func (b *BaseVCS) branchTitle(issue forge.Issue, prefix string, suffix string) string {
	title := b.formatTitle(issue, prefix, suffix)
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

	commitMessage := b.formatTitle(issue, branchPrefix, branchSuffix)
	b.impl.CreateEmptyCommit(commitMessage)
	b.impl.Push()
	return nil
}

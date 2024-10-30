package cmd

import (
	"gokid/forge"
	"gokid/shell"

	"github.com/spf13/cobra"
)

type Ready struct {
	forge forge.Forge
}

func NewReady(s shell.Shell) *Ready {
	return &Ready{
		forge: forge.NewGitHub(s),
	}
}

func (r *Ready) markReady() {
	r.forge.MarkPullRequestReady()
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "ready",
		Short: "Mark a change as ready for review",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			ready := NewReady(shell.New())
			ready.markReady()
		},
		Aliases: []string{"r"},
	})
}

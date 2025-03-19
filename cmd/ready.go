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

func (r *Ready) markReady() error {
	return r.forge.MarkPullRequestReady()
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "ready",
		Short: "Mark a change as ready for review",
		Long:  "",
		Run: func(_ *cobra.Command, _ []string) {
			ready := NewReady(shell.New())
			err := ready.markReady()
			if err != nil {
				panic(err)
			}
		},
		Aliases: []string{"r"},
	})
}

package cmd

import (
	"gokid/shell"

	"github.com/spf13/cobra"
)

type Ready struct {
	shell shell.Shell
}

func NewReady(s shell.Shell) *Ready {
	return &Ready{
		shell: s,
	}
}

func (r *Ready) markReady() {
	r.shell.Run("gh pr ready")
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

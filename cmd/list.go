package cmd

import (
	"fmt"
	"gokid/forge"
	"gokid/shell"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	var commitChanges bool

	cmd := &cobra.Command{
		Use:   "list [author]",
		Short: "List current PRs",
		Long:  "",
		Run: func(_ *cobra.Command, args []string) {
			shell := shell.New()

			var authors string
			switch len(args) {
			case 0:
				authors = "all"
			case 1:
				authors = args[0]
				if authors != "me" && authors != "all" {
					if authors == "m" {
						authors = "me"
					} else {
						fmt.Fprintf(os.Stderr, "invalid argument: %s\n", authors)
						os.Exit(1)
					}
				}
			default:
				fmt.Fprintf(os.Stderr, "too many arguments\n")
				os.Exit(1)
			}

			github := forge.NewGitHub(shell)
			github.ListPullRequests(authors)
		},
		Aliases: []string{"l"},
	}

	cmd.Flags().BoolVar(&commitChanges, "commit", false, "Commit current changes instead of stashing")
	rootCmd.AddCommand(cmd)
}

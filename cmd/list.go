package cmd

import (
	"fmt"
	"strings"

	"gokid/forge"
	"gokid/shell"

	"github.com/spf13/cobra"
)

// CIState is an enum representing overall CI status
type CIState int

const (
	CIUnknown CIState = iota
	CISuccess
	CIPending
	CIFailure
)

func humanCIState(state CIState) string {
	switch state {
	case CISuccess:
		return "✓ success"
	case CIPending:
		return "… pending"
	case CIFailure:
		return "✗ failing"
	default:
		return "?"
	}
}

func summarizeCI(pr forge.PullRequest) CIState {
	if len(pr.StatusCheckRollup) == 0 {
		return CIUnknown
	}
	// Determine worst status across items: any failure/error -> failure; any in-progress -> pending; else success if at least one success; else unknown
	hasInProgress := false
	hasSuccess := false
	for _, c := range pr.StatusCheckRollup {
		status := strings.ToLower(string(c.Status))
		if c.Conclusion == forge.CheckConclusionFailure || c.Conclusion == forge.CheckConclusionCancelled || c.Conclusion == forge.CheckConclusionTimedOut || c.Conclusion == forge.CheckConclusionActionRequired || c.Conclusion == forge.CheckConclusionError {
			return CIFailure
		}
		if status == "in_progress" || status == "queued" { // queued treated as pending
			hasInProgress = true
		}
		if c.Conclusion == forge.CheckConclusionSuccess {
			hasSuccess = true
		}
	}
	if hasInProgress {
		return CIPending
	}
	if hasSuccess {
		return CISuccess
	}
	return CIUnknown
}

const PRstr = "#%-5d  %-50s  %-10s  %s"

func renderPRRow(pr forge.PullRequest) string {
	ci := summarizeCI(pr)
	return fmt.Sprintf(PRstr, pr.Number, pr.Title, humanCIState(ci), pr.URL)
}

func listPRs(f forge.Forge) error {
	prs, err := f.ListPullRequests()
	if err != nil {
		return err
	}

	if len(prs) == 0 {
		fmt.Println("No open pull requests.")
		return nil
	}

	// Header
	fmt.Printf(PRstr, 0, "TITLE", "CI STATUS", "URL")
	fmt.Println()
	for _, pr := range prs {
		fmt.Println(renderPRRow(pr))
	}
	return nil
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "list",
		Short: "List open pull requests for this repository",
		Run: func(_ *cobra.Command, _ []string) {
			f := forge.NewGitHub(shell.New())
			if err := listPRs(f); err != nil {
				panic(err)
			}
		},
	})
}

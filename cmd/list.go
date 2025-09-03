package cmd

import (
	"fmt"
	"strings"

	"gokid/forge"
	"gokid/shell"

	"github.com/spf13/cobra"
)

func humanCIState(state string) string {
	switch strings.ToLower(state) {
	case "success":
		return "✓ success"
	case "pending", "expected":
		return "… pending"
	case "failure", "action_required", "cancelled", "error", "timed_out":
		return "✗ failing"
	default:
		if state == "" {
			return "?"
		}
		return state
	}
}

func summarizeCI(pr forge.PullRequest) string {
	if len(pr.StatusCheckRollup) == 0 {
		return ""
	}
	// Determine worst status across items: any failure/error -> failure; any in-progress -> pending; else success if at least one success; else unknown
	hasInProgress := false
	hasSuccess := false
	for _, c := range pr.StatusCheckRollup {
			status := strings.ToLower(c.Status)
		conclusion := strings.ToLower(c.Conclusion)
		if conclusion == "failure" || conclusion == "cancelled" || conclusion == "timed_out" || conclusion == "action_required" || conclusion == "error" {
			return "failure"
		}
		if status == "in_progress" || status == "queued" { // queued treated as pending
			hasInProgress = true
		}
		if conclusion == "success" {
			hasSuccess = true
		}
	}
	if hasInProgress {
		return "pending"
	}
	if hasSuccess {
		return "success"
	}
	return ""
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

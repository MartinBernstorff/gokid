package forge

import (
	"encoding/json"
	"strings"
)

// PRState is an enum for pull request state
type PRState string

const (
	PRStateUnknown PRState = "UNKNOWN"
	PRStateOpen    PRState = "OPEN"
	PRStateClosed  PRState = "CLOSED"
	PRStateMerged  PRState = "MERGED"
)

func (s *PRState) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	vv := strings.ToUpper(v)
	switch PRState(vv) {
	case PRStateOpen, PRStateClosed, PRStateMerged:
		*s = PRState(vv)
	default:
		*s = PRStateUnknown
	}
	return nil
}

// CheckStatus is an enum for CI check run status
type CheckStatus string

const (
	CheckStatusUnknown    CheckStatus = "UNKNOWN"
	CheckStatusQueued     CheckStatus = "QUEUED"
	CheckStatusInProgress CheckStatus = "IN_PROGRESS"
	CheckStatusCompleted  CheckStatus = "COMPLETED"
)

func (s *CheckStatus) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	vv := strings.ToUpper(v)
	switch CheckStatus(vv) {
	case CheckStatusQueued, CheckStatusInProgress, CheckStatusCompleted:
		*s = CheckStatus(vv)
	default:
		*s = CheckStatusUnknown
	}
	return nil
}

// CheckConclusion is an enum for CI check conclusion
type CheckConclusion string

const (
	CheckConclusionUnknown        CheckConclusion = "UNKNOWN"
	CheckConclusionSuccess        CheckConclusion = "SUCCESS"
	CheckConclusionFailure        CheckConclusion = "FAILURE"
	CheckConclusionNeutral        CheckConclusion = "NEUTRAL"
	CheckConclusionCancelled      CheckConclusion = "CANCELLED"
	CheckConclusionTimedOut       CheckConclusion = "TIMED_OUT"
	CheckConclusionActionRequired CheckConclusion = "ACTION_REQUIRED"
	CheckConclusionStale          CheckConclusion = "STALE"
	CheckConclusionSkipped        CheckConclusion = "SKIPPED"
	CheckConclusionStartupFailure CheckConclusion = "STARTUP_FAILURE"
	CheckConclusionError          CheckConclusion = "ERROR"
)

func (c *CheckConclusion) UnmarshalJSON(b []byte) error {
	// conclusion can be null/empty for in-progress checks; handle that
	var v *string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	if v == nil {
		*c = CheckConclusionUnknown
		return nil
	}
	vv := strings.ToUpper(*v)
	switch CheckConclusion(vv) {
	case CheckConclusionSuccess, CheckConclusionFailure, CheckConclusionNeutral, CheckConclusionCancelled,
		CheckConclusionTimedOut, CheckConclusionActionRequired, CheckConclusionStale, CheckConclusionSkipped,
		CheckConclusionStartupFailure, CheckConclusionError:
		*c = CheckConclusion(vv)
	default:
		*c = CheckConclusionUnknown
	}
	return nil
}

// CheckType is the GraphQL typename of the CI status entry
type CheckType string

const (
	CheckTypeUnknown       CheckType = "UNKNOWN"
	CheckTypeCheckRun      CheckType = "CHECKRUN"      // from "CheckRun"
	CheckTypeStatusContext CheckType = "STATUSCONTEXT" // from "StatusContext"
)

func (t *CheckType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	vv := strings.ToUpper(strings.ReplaceAll(v, "_", ""))
	switch CheckType(vv) {
	case CheckTypeCheckRun, CheckTypeStatusContext:
		*t = CheckType(vv)
	default:
		*t = CheckTypeUnknown
	}
	return nil
}

// CheckItem represents a single element in the statusCheckRollup array
type CheckItem struct {
	TypeName     CheckType       `json:"__typename"`
	Name         string          `json:"name"`
	Status       CheckStatus     `json:"status"`
	Conclusion   CheckConclusion `json:"conclusion"`
	WorkflowName string          `json:"workflowName"`
	DetailsURL   string          `json:"detailsUrl"`
}

// PullRequest models a PR with key fields used across implementations
type PullRequest struct {
	Number int     `json:"number,omitempty"`
	Title  string  `json:"title,omitempty"`
	State  PRState `json:"state,omitempty"`
	URL    string  `json:"url,omitempty"`
	Base   string  `json:"-"`
	Draft  bool    `json:"isDraft,omitempty"`

	// CI status information from GitHub GraphQL via gh --json statusCheckRollup
	StatusCheckRollup []CheckItem `json:"statusCheckRollup,omitempty"`
}

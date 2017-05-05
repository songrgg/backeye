package assertion

import (
	"context"
	nethttp "net/http"
	"time"
)

// Result records the assertion's result
type Result struct {
	AssertionID       int64         `json:"assertion_id"`
	Passed            bool          `json:"passed"`
	Error             error         `json:"error"`
	ExecutionDuration time.Duration `json:"execution_duration"`
}

// AssertionFunc indicates the assertion
type AssertionFunc func(context.Context, *nethttp.Response) Result

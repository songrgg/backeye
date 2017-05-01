package assertion

import (
	"context"
	nethttp "net/http"
	"time"
)

// AssertionResult records the assertion's result
type AssertionResult struct {
	AssertionID       int64
	Success           bool
	Error             error
	ExecutionDuration time.Duration
}

// AssertionFunc indicates the assertion
type AssertionFunc func(context.Context, *nethttp.Response) AssertionResult

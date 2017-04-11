package assertion

import (
	"context"
	nethttp "net/http"
)

// AssertionResult records the assertion's result
type AssertionResult struct {
	Success bool
	Error   error
}

// AssertionFunc indicates the assertion
type AssertionFunc func(context.Context, *nethttp.Response) AssertionResult

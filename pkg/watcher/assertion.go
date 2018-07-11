package watcher

import "github.com/songrgg/backeye/pkg/response"

// Assertion wraps the watcher assertion, including assertion operation, key, expected.
type Assertion interface {
	Check(res response.Response) (bool, error)
}

// AssertionResult wraps the result of the assertion.
type AssertionResult struct {
	Result bool
	Err    error
}

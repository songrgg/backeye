package http

import (
	"context"
	"errors"
	"io/ioutil"
	nethttp "net/http"
	"time"

	"github.com/songrgg/backeye/assertion"
	"github.com/songrgg/backeye/watch"
)

// Method indicates the HTTP methods
type Method int

// AssertionFunc indicates the assertion
// type AssertionFunc func(context.Context, *nethttp.Response) AssertionResult

const (
	// GET indicates the HTTP Get method
	GET Method = iota
	// POST indicates the HTTP Get method
	POST Method = iota
	// PUT indicates the HTTP Get method
	PUT Method = iota
	// DELETE indicates the HTTP Get method
	DELETE Method = iota
	// HEAD indicates the HTTP head method
	HEAD Method = iota
)

// Watch defines API sampling
type Watch struct {
	Name          string
	Desc          string
	Method        Method
	Path          string
	PathVariables map[string]string
	Headers       map[string]string
	Timeout       time.Duration
	ExtractFunc   func(string) map[string]string
	Assertions    []assertion.AssertionFunc
}

// Run executes this watch
func (w *Watch) Run(ctx context.Context) (watch.WatchResult, error) {
	switch w.Method {
	case GET:
		resp, err := nethttp.Get(w.Path)
		if err != nil {
			return watch.WatchResult{}, err
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return watch.WatchResult{}, err
		}

		ctx = context.WithValue(ctx, watch.ResponseBody, body)

		assertionResults := make([]assertion.AssertionResult, 0)
		for i, assertion := range w.Assertions {
			assertionResults = append(assertionResults, assertion(ctx, resp))
			if !assertionResults[i].Success {
				break
			}
		}
		return watch.WatchResult{
			Response:   resp,
			Assertions: assertionResults,
		}, nil

	case POST:
	case PUT:
		return watch.WatchResult{}, nil
	}

	return watch.WatchResult{}, errors.New("unsupported HTTP method")
}

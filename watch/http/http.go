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
func (w *Watch) Run(ctx context.Context) (watch.Result, error) {
	var taskID int64
	if ctx.Value("task_id") != nil {
		taskID = ctx.Value("task_id").(int64)
	} else {
		taskID = 0
	}
	result := watch.Result{
		TaskID:        taskID,
		WatchName:     w.Name,
		ExecutionTime: time.Now(),
	}
	switch w.Method {
	case GET:
		resp, err := nethttp.Get(w.Path)
		if err != nil {
			return result, err
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return result, err
		}

		ctx = context.WithValue(ctx, watch.ResponseBody, body)

		assertionResults := make([]assertion.AssertionResult, 0)
		for i, assertion := range w.Assertions {
			assertionResults = append(assertionResults, assertion(ctx, resp))
			if !assertionResults[i].Success {
				break
			}
		}
		result.Response = resp
		result.Assertions = assertionResults
		return result, nil

	case POST:
	case PUT:
		return result, nil
	}

	return result, errors.New("unsupported HTTP method")
}

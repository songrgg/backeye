package parser

import (
	"context"
	"encoding/json"
	"errors"
	nethttp "net/http"
	"strconv"
	"time"

	"github.com/songrgg/backeye/assertion"
	"github.com/songrgg/backeye/model"
	"github.com/songrgg/backeye/std"
	"github.com/songrgg/backeye/task"
	"github.com/songrgg/backeye/watch"
	"github.com/songrgg/backeye/watch/http"
	modelmapper "gopkg.in/jeevatkm/go-model.v0"
)

// DefaultParser represents the default parser
type DefaultParser struct {
}

// ParseTask translates model to task
func (p *DefaultParser) ParseTask(t *model.Task) (*task.Task, error) {
	task := task.Task{}
	modelmapper.Copy(&task, t)

	watches := make([]watch.Watch, 0)
	for i := range t.Watches {
		watch, err := p.ParseWatch(&t.Watches[i])
		if err != nil {
			std.LogErrorc("default_parser", err, "failed to parse watch")
		}
		watches = append(watches, watch)
	}
	task.Watches = watches
	return &task, nil
}

func (p *DefaultParser) ParseWatch(w *model.Watch) (watch.Watch, error) {
	newwatch := &http.Watch{
		Name: w.Name,
		Desc: w.Desc,
		Path: w.Path,
	}

	// parse method
	switch w.Method {
	case "GET":
		newwatch.Method = http.GET
	case "POST":
		newwatch.Method = http.POST
	case "PUT":
		newwatch.Method = http.PUT
	case "HEAD":
		newwatch.Method = http.HEAD
	}

	// parse path variables
	// pathVars := make(map[string]string)
	// for _, pathvar := range w.PathVariables {
	// 	pathVars[pathvar.Name] = pathvar.Value
	// }
	// watch.PathVariables = pathVars

	if w.Timeout > 0 {
		newwatch.Timeout = time.Duration(w.Timeout) * time.Millisecond
	}

	assertions := make([]assertion.AssertionFunc, 0)
	for i := range w.Assertions {
		assertions = append(assertions, p.parseAssertion(&w.Assertions[i]))
	}
	newwatch.Assertions = assertions

	variables := make([]watch.Variable, 0)
	for i := range w.Variables {
		v, err := p.parseVariable(&w.Variables[i])
		if err != nil {
			continue
		}
		variables = append(variables, v)
	}
	newwatch.Variables = variables

	return newwatch, nil
}

func (p *DefaultParser) parseVariable(t *model.Variable) (watch.Variable, error) {
	var newvar watch.Variable
	errors := modelmapper.Copy(&newvar, t)
	if len(errors) > 0 {
		return newvar, errors[0]
	}
	return newvar, nil
}

func (p *DefaultParser) parseAssertion(t *model.Assertion) assertion.AssertionFunc {
	return func(ctx context.Context, resp *nethttp.Response) assertion.Result {
		start := time.Now()
		body := ctx.Value(watch.ResponseBody)

		v := make(map[string]interface{})
		err := json.Unmarshal(body.([]byte), &v)
		if err != nil {
			return assertion.Result{
				Success: false,
				Error:   err,
			}
		}

		left := ""
		if t.Source == "header" {
			if t.Left == "status_code" {
				left = strconv.Itoa(resp.StatusCode)
			} else {
				left = resp.Header.Get(t.Left)
			}
		} else if t.Source == "body" {
			switch v[t.Left].(type) {
			case float64:
				left = strconv.FormatFloat(v[t.Left].(float64), 'g', -1, 64)
			default:
				left = v[t.Left].(string)
			}
		} else {
			return assertion.Result{
				Success: false,
				Error:   errors.New("invalid source"),
			}
		}

		right := t.Right

		success := false

		if t.Operator == "equal" {
			success = left == right
			err = errors.New("not equal")
		} else if t.Operator == "not_empty" {
			success = left != ""
			err = errors.New("not empty")
		}
		return assertion.Result{
			AssertionID:       t.ID,
			ExecutionDuration: time.Since(start),
			Success:           success,
			Error:             err,
		}
	}
}

// Translate translates JSON to task
func (p *DefaultParser) Translate(data interface{}) (*task.Task, error) {
	return p.ParseTask(data.(*model.Task))
}

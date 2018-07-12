package watcher

import (
	"context"
	"errors"
	"fmt"

	"github.com/songrgg/backeye/pkg/common"
	"github.com/songrgg/backeye/pkg/task"
	"github.com/songrgg/backeye/pkg/variable"
)

var (
	ErrAssertFailure = errors.New("failed to pass assertions")
)

// Point wraps the watcher point, like watch task, assertions, variables extraction.
type Point struct {
	Task       task.Task
	Assertions []Assertion
	Variables  []variable.Variable
}

func (p *Point) Do(ctx context.Context) (context.Context, error) {
	res, err := p.Task.Do(ctx)

	ctx2 := context.WithValue(ctx, common.ResponseKey, res)
	if err != nil {
		return ctx2, err
	}

	var assertionResults []AssertionResult
	for _, assert := range p.Assertions {
		result, err := assert.Check(res)
		assertionResults = append(assertionResults, AssertionResult{Result: result, Err: err})
		if err != nil {
			return context.WithValue(ctx2, common.AssertionResultsKey, assertionResults), ErrAssertFailure
		}
	}

	ctx2 = context.WithValue(ctx2, common.AssertionResultsKey, assertionResults)

	// extract variables to the context
	variables := make(map[string]interface{})
	for _, v := range p.Variables {
		key, val, err := v.Extract(res)
		if err != nil {
			fmt.Println("fail to extract value")
			continue
		}
		variables[key] = val
	}
	ctx3 := context.WithValue(ctx2, common.VariablesKey, variables)

	return ctx3, nil
}

func (p *Point) AddAssertion(a Assertion) {
	if p.Assertions == nil {
		p.Assertions = make([]Assertion, 0)
	}
	p.Assertions = append(p.Assertions, a)
}

func (p *Point) AddVariable(v variable.Variable) {
	if p.Variables == nil {
		p.Variables = make([]variable.Variable, 0)
	}
	p.Variables = append(p.Variables, v)
}

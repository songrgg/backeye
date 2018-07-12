package watcher

import (
	"context"
	"testing"

	"github.com/songrgg/backeye/pkg/common"
	"github.com/songrgg/backeye/pkg/response"
	"github.com/songrgg/backeye/pkg/watcher/http"
	"github.com/stretchr/testify/assert"
)

func TestPassPoint(t *testing.T) {
	point := Point{
		Task: &http.Task{
			Path:    "http://httpbin.org/get",
			Method:  http.Get,
			Timeout: 0,
		},
		Assertions: nil,
		Variables:  nil,
	}

	assert.NotNil(t, point.Task)
	assert.Nil(t, point.Assertions)
	assert.Nil(t, point.Variables)

	statusAssertion, err := http.NewAssertion("equal", http.SourceStatus, "", 200, "")
	assert.Nil(t, err)
	point.AddAssertion(statusAssertion)
	assert.NotNil(t, point.Assertions)

	ctx2, err := point.Do(context.Background())
	assert.NotNil(t, ctx2.Value(common.ResponseKey))
	assert.NotNil(t, ctx2.Value(common.AssertionResultsKey))
	assert.NotNil(t, ctx2.Value(common.VariablesKey))
	assert.Empty(t, ctx2.Value(common.VariablesKey))

	res := ctx2.Value(common.ResponseKey).(response.Response)
	assert.NotNil(t, http.GetHeaders(res))
	assert.NotNil(t, http.GetStatus(res))
	assert.NotNil(t, http.GetBody(res))

	assertionResults := ctx2.Value(common.AssertionResultsKey).([]AssertionResult)
	assert.Len(t, assertionResults, 1)
	assert.True(t, assertionResults[0].Result)
}

func TestFailPoint(t *testing.T) {
	point := Point{
		Task: &http.Task{
			Path:    "http://httpbin.org/get",
			Method:  http.Get,
			Timeout: 0,
		},
		Assertions: nil,
		Variables:  nil,
	}

	assert.NotNil(t, point.Task)
	assert.Nil(t, point.Assertions)
	assert.Nil(t, point.Variables)

	statusAssertion, err := http.NewAssertion("equal", http.SourceStatus, "", 200, "")
	assert.Nil(t, err)
	point.AddAssertion(statusAssertion)
	assert.NotNil(t, point.Assertions)

	failAssertion, err := http.NewAssertion("equal", http.SourceStatus, "", 500, "")
	assert.Nil(t, err)
	point.AddAssertion(failAssertion)

	ctx2, err := point.Do(context.Background())
	assert.NotNil(t, ctx2.Value(common.ResponseKey))
	assert.NotNil(t, ctx2.Value(common.AssertionResultsKey))
	assert.Empty(t, ctx2.Value(common.VariablesKey))

	res := ctx2.Value(common.ResponseKey).(response.Response)
	assert.NotNil(t, http.GetHeaders(res))
	assert.NotNil(t, http.GetStatus(res))
	assert.NotNil(t, http.GetBody(res))

	assertionResults := ctx2.Value(common.AssertionResultsKey).([]AssertionResult)
	assert.Len(t, assertionResults, 2)
	assert.True(t, assertionResults[0].Result)
	assert.False(t, assertionResults[1].Result)
}

func TestVariableExtract(t *testing.T) {
	point := Point{
		Task: &http.Task{
			Path:    "http://httpbin.org/get",
			Method:  http.Get,
			Timeout: 0,
		},
		Assertions: nil,
		Variables:  nil,
	}

	point.AddVariable(&http.Variable{
		Source:         http.SourceBody,
		Key:            "url",
		SourceEncoding: "json",
	})

	assert.Len(t, point.Variables, 1)

	ctx, err := point.Do(context.Background())
	assert.Nil(t, err)

	vars := ctx.Value(common.VariablesKey)
	assert.NotNil(t, vars)

	varsMap := vars.(map[string]interface{})
	assert.Len(t, varsMap, 1)
	assert.Equal(t, "http://httpbin.org/get", varsMap["url"])
}

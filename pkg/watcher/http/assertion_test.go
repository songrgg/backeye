package http

import (
	"testing"

	"net/http"

	"github.com/songrgg/backeye/pkg/response"
	"github.com/stretchr/testify/assert"
)

func TestUnavailable(t *testing.T) {
	a, err := NewAssertion("xxx", "", "", "", "")
	assert.NotNil(t, err)
	assert.Nil(t, a)
}

func TestCheckHeader(t *testing.T) {
	a, err := NewAssertion("equal", "header", "X-Upstream-Status", "200", "")
	assert.Nil(t, err)

	res := response.Response{}

	result, err := a.Check(res)
	assert.NotNil(t, err)
	assert.False(t, result)

	SetHeaders(res, http.Header{
		"X-Upstream-Status": []string{"200"},
	})
	result, err = a.Check(res)
	assert.Nil(t, err)
	assert.True(t, result)
}

func TestCheckStatus(t *testing.T) {
	a, err := NewAssertion("equal", "status", "", 200, "")
	assert.Nil(t, err)

	res := response.Response{}

	result, err := a.Check(res)
	assert.NotNil(t, err)
	assert.False(t, result)

	SetStatus(res, 500)
	result, err = a.Check(res)
	assert.NotNil(t, err)
	assert.False(t, result)

	SetStatus(res, 200)
	result, err = a.Check(res)
	assert.Nil(t, err)
	assert.True(t, result)
}

func TestCheckBody(t *testing.T) {
	a, err := NewAssertion("equal", "body", "code", float64(20000), "json")
	assert.Nil(t, err)

	res := response.Response{}

	result, err := a.Check(res)
	assert.NotNil(t, err)
	assert.False(t, result)

	SetBody(res, []byte("{\"code\": 50000}"))
	result, err = a.Check(res)
	assert.NotNil(t, err)
	assert.False(t, result)

	SetBody(res, []byte("{\"code\": 20000}"))
	result, err = a.Check(res)
	assert.Nil(t, err)
	assert.True(t, result)
}

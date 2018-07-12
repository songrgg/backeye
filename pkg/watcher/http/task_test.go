package http

import (
	"context"
	"testing"

	"encoding/json"

	"github.com/songrgg/backeye/pkg/common"
	"github.com/stretchr/testify/assert"
)

func TestNewHTTPTask(t *testing.T) {
	task := Task{
		Path:    "http://httpbin.org/get",
		Method:  Get,
		Timeout: 0,
	}
	res, err := task.Do(context.Background())
	body := GetBody(res)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.NotNil(t, body)
	assert.NotNil(t, GetHeaders(res))
	assert.Equal(t, 200, GetStatus(res))
	assert.Equal(t, "application/json", GetHeader(res, "content-type"))

	var jsonBody map[string]interface{}
	json.Unmarshal(body, &jsonBody)
	assert.Equal(t, "http://httpbin.org/get", jsonBody["url"])
}

func TestVariable(t *testing.T) {
	task := Task{
		Path:    "http://httpbin.org/{{method}}",
		Method:  Get,
		Timeout: 0,
	}
	ctx := context.Background()
	ctx = context.WithValue(ctx, common.VariablesKey, map[string]interface{}{
		"method": "get",
	})
	res, err := task.Do(ctx)
	body := GetBody(res)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.NotNil(t, body)
	assert.NotNil(t, GetHeaders(res))
	assert.Equal(t, 200, GetStatus(res))

	var jsonBody map[string]interface{}
	json.Unmarshal(body, &jsonBody)
	assert.Equal(t, "http://httpbin.org/get", jsonBody["url"])
}

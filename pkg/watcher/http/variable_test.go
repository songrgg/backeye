package http

import (
	"net/http"
	"testing"

	"github.com/songrgg/backeye/pkg/response"
	"github.com/stretchr/testify/assert"
)

func TestExtractHeader(t *testing.T) {
	headerContentType := "Content-Type"
	valueContentType := "application/json"
	v := &Variable{
		Source:         SourceHeader,
		Key:            headerContentType,
		SourceEncoding: "",
	}

	res := response.Response{}
	SetHeaders(res, http.Header{
		headerContentType: []string{
			valueContentType,
		},
	})

	key, val, err := v.Extract(res)
	assert.Nil(t, err)
	assert.Equal(t, headerContentType, key)
	assert.Equal(t, valueContentType, val)
}

func TestExtractBody(t *testing.T) {
	v := &Variable{
		Source:         SourceBody,
		Key:            "id",
		SourceEncoding: "json",
	}

	res := response.Response{}
	SetBody(res, []byte("{\"id\":1000}"))

	key, val, err := v.Extract(res)
	assert.Nil(t, err)
	assert.Equal(t, "id", key)
	assert.Equal(t, float64(1000), val)
}

func TestExtractFailure(t *testing.T) {
	v := &Variable{
		Source:         "unknown_source",
		Key:            "id",
		SourceEncoding: "json",
	}

	res := response.Response{}
	SetBody(res, []byte("{\"id\":1000}"))

	key, val, err := v.Extract(res)
	assert.Equal(t, "id", key)
	assert.Nil(t, val)
	assert.NotNil(t, err)

	// Test not json body case
	v = &Variable{
		Source:         "body",
		Key:            "id",
		SourceEncoding: "not_json",
	}

	res = response.Response{}
	SetBody(res, []byte("{\"id\":1000}"))

	key, val, err = v.Extract(res)
	assert.Equal(t, "id", key)
	assert.Nil(t, val)
	assert.NotNil(t, err)
}

package watcher

import (
	"testing"

	"context"

	"github.com/songrgg/backeye/pkg/variable"
	"github.com/songrgg/backeye/pkg/watcher/http"
	"github.com/stretchr/testify/assert"
)

func TestWatcher(t *testing.T) {
	w := &Watcher{}
	url := "http://httpbin.org/get"
	as, _ := http.NewAssertion("equal", http.SourceBody, "url", url, "json")
	w.AddPoint(Point{
		Task: &http.Task{
			Path:   url,
			Method: http.Get,
		},
		Assertions: []Assertion{
			as,
		},
		Variables: []variable.Variable{
			&http.Variable{
				Source:         http.SourceBody,
				Key:            "url",
				SourceEncoding: "json",
			},
		},
	})
	w.AddPoint(Point{
		Task: &http.Task{
			Path:   "{{url}}",
			Method: http.Get,
		},
		Assertions: []Assertion{
			as,
		},
	})

	assertionResults, err := w.Do(context.Background())
	assert.Nil(t, err)
	assert.Len(t, assertionResults, 2)
}

func TestNewWatcher(t *testing.T) {
	conf := `
[
	{
		"name": "post list",
		"desc": "post list",
		"type": "http",
		"conf": "\n{\n    \"task\": {\n        \"path\": \"https://api-prod.wallstreetcn.com/apiv1/content/articles\",\n        \"method\": \"get\"\n    },\n    \"assertions\": [\n        {\n            \"source\": \"status\",\n            \"operator\": \"equal\",\n            \"expected\": 200\n        },\n        {\n            \"source\": \"body\",\n            \"operator\": \"equal\",\n            \"key\": \"code\",\n            \"expected\": 20000,\n            \"source_encoding\": \"json\"\n        }\n    ],\n    \"variables\": [\n        {\n            \"source\": \"body\",\n            \"key\": \"code\",\n            \"source_encoding\": \"json\"\n        }\n    ]\n}\n"
	}
]
`
	w, err := NewWatcher(conf)
	assert.Nil(t, err)
	assert.NotNil(t, w)

	results, err := w.Do(context.Background())
	assert.Nil(t, err)
	assert.Len(t, results, 2)
}

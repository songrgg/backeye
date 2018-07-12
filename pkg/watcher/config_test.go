package watcher

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConf(t *testing.T) {
	watcherConf := `{
    	"points": [
        	{
            	"name": "post list",
            	"desc": "post list",
            	"type": "http",
            	"conf": "\n{\n    \"task\": {\n        \"path\": \"https://api-prod.wallstreetcn.com/apiv1/content/articles\",\n        \"method\": \"get\"\n    },\n    \"assertions\": [\n        {\n            \"source\": \"status\",\n            \"operator\": \"equal\",\n            \"value\": 200\n        },\n        {\n            \"source\": \"body\",\n            \"operator\": \"equal\",\n            \"key\": \"code\",\n            \"value\": 20000,\n            \"source_encoding\": \"json\"\n        }\n    ],\n    \"variables\": [\n        {\n            \"source\": \"body\",\n            \"key\": \"code\",\n            \"source_encoding\": \"json\"\n        }\n    ]\n}\n"
        	}
    	]
	}
`
	c := Config{}
	json.Unmarshal([]byte(watcherConf), &c)
	assert.NotNil(t, c.Points)
	assert.Len(t, c.Points, 1)
}

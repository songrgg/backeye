package dao

import (
	"encoding/json"
	"testing"

	"github.com/songrgg/backeye/model/form"
)

// NewProject creates a new project
func TestNewWatcher(t *testing.T) {
	conf := `
{
    "name": "Post API",
    "desc": "post API",
    "cron": "*/2 * * * *",
    "points": [
        {
            "name": "post list",
            "desc": "post list",
            "type": "http",
            "conf": "\n{\n    \"task\": {\n        \"path\": \"https://api-prod.wallstreetcn.com/apiv1/content/articles\",\n        \"method\": \"get\"\n    },\n    \"assertions\": [\n        {\n            \"source\": \"status\",\n            \"operator\": \"equal\",\n            \"expected\": 200\n        },\n        {\n            \"source\": \"body\",\n            \"operator\": \"equal\",\n            \"key\": \"code\",\n            \"expected\": 20000,\n            \"source_encoding\": \"json\"\n        }\n    ],\n    \"variables\": [\n        {\n            \"source\": \"body\",\n            \"key\": \"code\",\n            \"source_encoding\": \"json\"\n        }\n    ]\n}\n"
        }
    ]
}
`
	w := form.Watcher{}
	json.Unmarshal([]byte(conf), &w)
	NewWatcher(&w)
}

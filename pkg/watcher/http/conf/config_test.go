package conf

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConf(t *testing.T) {
	confStr := `
{
    "task": {
        "path": "http://httpbin.org/get",
        "method": "get"
    },
    "assertions": [
        {
            "source": "status",
            "operator": "equal",
            "expected": 200
        },
        {
            "source": "body",
            "operator": "equal",
            "key": "url",
            "expected": "http://httpbin.org/get",
            "source_encoding": "json"
        }
    ],
    "variables": [
        {
            "source": "body",
            "key": "url",
            "source_encoding": "json"
        }
    ]
}
`
	conf := Config{}
	json.Unmarshal([]byte(confStr), &conf)
	assert.NotNil(t, conf.Task)
	assert.NotNil(t, conf.Variables)
	assert.NotNil(t, conf.Assertions)
}

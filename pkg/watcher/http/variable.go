package http

import (
	"encoding/json"
	"errors"

	"github.com/songrgg/backeye/pkg/response"
)

// Variable extracts the specified value from point response and register them in the variables section.
type Variable struct {
	Source         string `json:"source"`
	Key            string `json:"key"`
	SourceEncoding string `json:"source_encoding"`
}

func (v *Variable) Extract(res response.Response) (string, interface{}, error) {
	switch v.Source {
	case SourceHeader:
		// fetch from header
		return v.Key, GetHeader(res, v.Key), nil
	case SourceBody:
		// fetch from json body
		if v.SourceEncoding == "json" {
			var body map[string]interface{}
			// TODO support n-level field
			json.Unmarshal(GetBody(res), &body)
			return v.Key, body[v.Key], nil
		}
		return v.Key, nil, errors.New("unsupported source encoding")
	}
	return v.Key, nil, errors.New("variable not found")
}

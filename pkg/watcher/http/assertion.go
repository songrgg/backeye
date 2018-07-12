package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"

	"github.com/songrgg/backeye/pkg/assertion"
	"github.com/songrgg/backeye/pkg/response"
)

// Assertion wraps the watcher assertion, including assertion operation, key, expected.
type Assertion struct {
	op             assertion.Operator
	source         string
	key            string
	expected       interface{}
	sourceEncoding string
}

type AssertionConf struct {
	Source         string      `json:"source"`
	Operator       string      `json:"operator"`
	Key            string      `json:"key"`
	Expected       interface{} `json:"expected"`
	SourceEncoding string      `json:"source_encoding"`
}

func NewAssertionWithConf(conf *AssertionConf) (*Assertion, error) {
	return NewAssertion(conf.Operator, conf.Source, conf.Key, conf.Expected, conf.SourceEncoding)
}

// NewAssertion created an assertion with key and expected value.
func NewAssertion(op string, source string, key string, expected interface{}, sourceEncoding string) (*Assertion, error) {
	opFunc, err := assertion.NewOperator(op)
	if err != nil {
		return nil, err
	}
	return &Assertion{
		op:             opFunc,
		source:         source,
		key:            key,
		expected:       expected,
		sourceEncoding: sourceEncoding,
	}, nil
}

func (a *Assertion) Check(res response.Response) (bool, error) {
	switch a.source {
	case SourceHeader:
		return a.op(GetHeader(res, a.key), a.expected)

	case SourceBody:
		if a.sourceEncoding == "json" {
			var body map[string]interface{}
			// TODO support n-level field
			err := json.Unmarshal(GetBody(res), &body)
			if err != nil {
				return false, errors.New(fmt.Sprintf("fail to parse json body: %v", err))
			}
			return a.op(body[a.key], a.expected)
		}
		return false, errors.New("unsupported source encoding")

	case SourceStatus:
		status := GetStatus(res)

		typ := reflect.TypeOf(a.expected).Kind()
		if typ != reflect.Int {
			if typ == reflect.Float64 {
				return a.op(status, int(a.expected.(float64)))
			} else if typ == reflect.String {
				expected, err := strconv.ParseInt(a.expected.(string), 10, 32)
				if err != nil {
					return false, errors.New(fmt.Sprintf("fail to parse given status code: %v", err))
				}
				return a.op(status, int(expected))
			} else {
				return false, errors.New(fmt.Sprintf("unsupported status field type: %v", typ))
			}
		}

		return a.op(status, a.expected)
	}
	return false, errors.New("unsupported source")
}

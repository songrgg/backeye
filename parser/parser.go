package parser

import (
	"github.com/songrgg/backeye/parser/json"
	"github.com/songrgg/backeye/target"
)

// Parser translates the source to the api test target
type Parser interface {
	// Translate execute the input
	Translate(data []byte) (*target.Target, error)
}

const (
	JSON string = "json"
)

// NewParser creates a new parser according to the target type
func NewParser(targetType string, input string) Parser {
	switch targetType {
	case JSON:
		return &json.Parser{}
	}
	return nil
}

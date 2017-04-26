package parser

import (
	"github.com/songrgg/backeye/parser/json"
	"github.com/songrgg/backeye/task"
)

// Parser translates the source to the api test task
type Parser interface {
	// Translate execute the input
	Translate(data interface{}) (*task.Task, error)
}

const (
	JSON string = "json"
)

// NewParser creates a new parser according to the task type
func NewParser(taskType string, input string) Parser {
	switch taskType {
	case JSON:
		return &json.Parser{}
	default:
		return &DefaultParser{}
	}
}

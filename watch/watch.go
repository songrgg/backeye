package watch

import (
	"context"
	"time"

	"github.com/songrgg/backeye/assertion"
)

type ResponseKey string

const (
	ResponseBody = ResponseKey("response_key")
)

// Watch defines an API watch
type Watch interface {
	Run(ctx context.Context) (Result, error)
}

// Result indicates the watch's result
type Result struct {
	TaskID        int64              `json:"task_id"`
	TaskName      string             `json:"task_name"`
	WatchName     string             `json:"watch_name"`
	ExecutionTime time.Time          `json:"execution_time"`
	ExtractValues map[string]string  `json:"extract_values"`
	Assertions    []assertion.Result `json:"assertions"`
}

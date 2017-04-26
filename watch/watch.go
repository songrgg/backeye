package watch

import (
	"context"
	nethttp "net/http"
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
	TaskID        int64
	TaskName      string
	WatchName     string
	ExecutionTime time.Time
	ExtractValues map[string]string
	Response      *nethttp.Response
	Assertions    []assertion.AssertionResult
}

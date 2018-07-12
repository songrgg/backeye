package task

import (
	"context"

	"github.com/songrgg/backeye/pkg/response"
)

// Task abstracts the watcher's task, like HTTP fetch, GRPC fetch, shell scripts.
// Fill the response with the task results.
type Task interface {
	Do(ctx context.Context) (response.Response, error)
}

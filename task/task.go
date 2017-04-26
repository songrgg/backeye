package task

import (
	"context"

	"github.com/songrgg/backeye/watch"
)

// Task wraps the whole API task test
type Task struct {
	ID       int64
	Name     string
	Desc     string
	CronSpec string
	Watches  []watch.Watch
}

// Run executes the watches in order
func (t *Task) Run(ctx context.Context) ([]watch.Result, error) {
	watchResults := make([]watch.Result, 0)
	ctx = context.WithValue(ctx, "task_id", t.ID)
	for _, watch := range t.Watches {
		watchResult, err := watch.Run(ctx)
		if err != nil {
			return watchResults, err
		}
		watchResults = append(watchResults, watchResult)
	}
	return watchResults, nil
}

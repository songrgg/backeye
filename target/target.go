package target

import (
	"context"

	"github.com/songrgg/backeye/watch"
)

// Target wraps the whole API target test
type Target struct {
	Name     string
	Desc     string
	CronSpec string
	Watches  []watch.Watch
}

// Run executes the watches in order
func (t *Target) Run(ctx context.Context) ([]watch.WatchResult, error) {
	watchResults := make([]watch.WatchResult, 0)
	for _, watch := range t.Watches {
		watchResult, err := watch.Run(ctx)
		watchResults = append(watchResults, watchResult)
		if err != nil {
			return watchResults, err
		}
	}
	return watchResults, nil
}

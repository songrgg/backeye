package form

import "github.com/songrgg/backeye/pkg/watcher"

type Watcher struct {
	Name   string                `json:"name"`
	Desc   string                `json:"desc"`
	Cron   string                `json:"cron"`
	Points []watcher.PointConfig `json:"points"`
}

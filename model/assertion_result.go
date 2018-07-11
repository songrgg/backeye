package model

import "github.com/songrgg/backeye/std"

type AssertionResult struct {
	ID              int64  `gorm:"primary_key" json:"id"`
	WatcherID       int64  `gorm:"index" json:"watcher_id"`
	WatcherRevision string `json:"watcher_revision"`
	Passed          bool   `json:"passed"`
	Result          string `gorm:"type:text"`
	std.TimeMixin
}

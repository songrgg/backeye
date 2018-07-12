package model

import "github.com/songrgg/backeye/std"

type Watcher struct {
	ID       int64  `gorm:"primary_key" json:"id"`
	Name     string `gorm:"type:varchar(64)" json:"name"`
	Desc     string `gorm:"type:varchar(512)" json:"desc"`
	Cron     string `gorm:"type:varchar(64)" json:"cron"`
	Points   string `gorm:"type:text" json:"points"`
	Disabled bool   `gorm:"type:boolean" json:"disabled"`
	std.TimeMixin
}

type WatcherSnapshot struct {
	ID       int64  `json:"id"`
	Revision int64  `json:"revision"`
	Name     string `gorm:"type:varchar(64)" json:"name"`
	Desc     string `gorm:"type:varchar(512)" json:"desc"`
	Cron     string `gorm:"type:varchar(64)" json:"cron"`
	Points   string `gorm:"type:text" json:"points"`
	std.TimeMixin
}

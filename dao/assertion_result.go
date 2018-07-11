package dao

import (
	"time"

	"fmt"

	"github.com/songrgg/backeye/model"
	"github.com/songrgg/backeye/std"
)

// NewAssertionResult creates a assertion result for watcher
func NewAssertionResult(watcherId int64, passed bool, result []byte) error {
	// fetch the greatest revision
	snapshot := model.WatcherSnapshot{}
	db := model.DB().First(&snapshot).Where("id = ?", watcherId).Order("revision desc")
	if db.Error != nil {
		return db.Error
	}

	nowTime := time.Now()
	db = model.DB().Save(&model.AssertionResult{
		WatcherID:       snapshot.ID,
		WatcherRevision: snapshot.Revision,
		Passed:          passed,
		Result:          string(result),
		TimeMixin: std.TimeMixin{
			CreatedAt: nowTime,
			UpdatedAt: nowTime,
		},
	})
	if db.Error != nil {
		fmt.Println(db.Error)
		return db.Error
	}

	return nil
}

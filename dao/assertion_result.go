package dao

import (
	"fmt"

	"github.com/songrgg/backeye/model"
)

// NewAssertionResult creates a assertion result for watcher
func NewAssertionResult(watcherId int64, passed bool, result []byte) error {
	watcherRevision := 0
	w := model.WatcherSnapshot{}
	db := model.DB().Find(&w, "id = ? and revision = ?", watcherId, watcherRevision)
	if db.Error != nil {
		return db.Error
	}

	fmt.Println(w)

	return nil
}

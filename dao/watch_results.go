package dao

import "github.com/songrgg/backeye/model"

// NewWatchResult creates a new watch result
func NewWatchResult(watchResult *model.WatchResult) error {
	db := model.DB().Save(watchResult)
	if err := db.Error; err != nil {
		return err
	}
	return nil
}

// AllWatchResults fetches all watch results
func AllWatchResults(taskName string, maxID string, limit int) ([]model.WatchResult, error) {
	var results []model.WatchResult
	db := model.DB().Find(&results)
	if err := db.Error; err != nil {
		return nil, err
	}
	return results, nil
}

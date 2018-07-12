package dao

import (
	"fmt"
	"time"

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

// GetAssertionResults fetches the specified watchID's assertion results.
func GetAssertionResults(watchID int64, orderBy string, limit int64) ([]model.AssertionResult, error) {
	if orderBy == "" {
		orderBy = "created_at desc"
	}

	rows, err := model.DB().Model(&model.AssertionResult{}).Where(`watcher_id = ?`, watchID).Order(orderBy).Limit(limit).Rows()
	defer rows.Close()

	if err != nil {
		return nil, err
	}
	var results []model.AssertionResult
	for rows.Next() {
		var result model.AssertionResult
		model.DB().ScanRows(rows, &result)
		results = append(results, result)
	}
	return results, nil
}

// GetLatestAssertionResults fetches the latest watchID's assertion results.
func GetLatestAssertionResults(watchIDs []int64) ([]model.AssertionResult, error) {
	rows, err := model.DB().Model(&model.AssertionResult{}).Where(`id IN (
	SELECT MAX(id) result_id FROM assertion_results WHERE watcher_id in (?) GROUP BY watcher_id
)
`, watchIDs).Rows()

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	var results []model.AssertionResult
	for rows.Next() {
		var result model.AssertionResult
		model.DB().ScanRows(rows, &result)
		results = append(results, result)
	}
	return results, nil
}

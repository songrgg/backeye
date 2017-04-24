package dao

import (
	"github.com/songrgg/backeye/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// NewWatchResult creates a new watch result
func NewWatchResult(watchResult *model.WatchResult) error {
	index := mgo.Index{
		Key:        []string{"task_name", "watch_name", "execution_time"},
		Unique:     true,
		DropDups:   true,
		Background: true, // See notes.
		Sparse:     true,
	}
	err := watchResultCollection.EnsureIndex(index)
	if err != nil && !mgo.IsDup(err) {
		return err
	}

	err = watchResultCollection.Insert(watchResult)
	if err != nil {
		if mgo.IsDup(err) {
			return nil
		}
		return err
	}
	return nil
}

// AllWatchResults fetches all watch results
func AllWatchResults(taskName string, maxID string, limit int) ([]model.WatchResult, error) {
	cnt, err := watchResultCollection.Count()
	if err != nil {
		return nil, err
	}

	cond := bson.M{
		"taskname": taskName,
	}
	if maxID != "" {
		cond["_id"] = bson.M{
			"$lt": maxID,
		}
	}

	allWatchResults := make([]model.WatchResult, cnt)
	iter := watchResultCollection.Find(cond).Sort("-_id").Limit(limit).Iter()
	if err := iter.All(&allWatchResults); err != nil {
		return nil, err
	}
	return allWatchResults, nil
}
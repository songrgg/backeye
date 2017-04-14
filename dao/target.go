package dao

import (
	"github.com/songrgg/backeye/model"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func NewTask(target *model.Task) error {
	index := mgo.Index{
		Key:        []string{"name"},
		Unique:     true,
		DropDups:   true,
		Background: true, // See notes.
		Sparse:     true,
	}
	err := targetCollection.EnsureIndex(index)
	if err != nil && !mgo.IsDup(err) {
		return err
	}

	err = targetCollection.Insert(target)
	if err != nil {
		if mgo.IsDup(err) {
			return nil
		}
		return err
	}
	return nil
}

func GetTask(name string) (*model.Task, error) {
	target := model.Task{}
	err := targetCollection.Find(bson.M{"name": name}).One(&target)
	return &target, err
}

func RemoveTask(name string) error {
	return targetCollection.Remove(bson.M{"name": name})
}

func UpdateTask(target *model.Task) error {
	return targetCollection.Update(bson.M{"name": target.Name}, target)
}

func AllTasks() ([]model.Task, error) {
	cnt, err := targetCollection.Count()
	if err != nil {
		return nil, err
	}

	allTasks := make([]model.Task, cnt)
	iter := targetCollection.Find(nil).Limit(1000).Iter()
	if err := iter.All(&allTasks); err != nil {
		return nil, err
	}
	return allTasks, nil
}

func ListTask(page int, limit int) ([]*model.Task, error) {
	return nil, nil
}

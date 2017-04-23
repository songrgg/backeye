package dao

import (
	"github.com/songrgg/backeye/model"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func NewTask(task *model.Task) error {
	index := mgo.Index{
		Key:        []string{"name"},
		Unique:     true,
		DropDups:   true,
		Background: true, // See notes.
		Sparse:     true,
	}
	err := taskCollection.EnsureIndex(index)
	if err != nil && !mgo.IsDup(err) {
		return err
	}

	err = taskCollection.Insert(task)
	if err != nil {
		if mgo.IsDup(err) {
			return nil
		}
		return err
	}
	return nil
}

func GetTask(name string) (*model.Task, error) {
	task := model.Task{}
	err := taskCollection.Find(bson.M{"name": name}).One(&task)
	return &task, err
}

func RemoveTask(name string) error {
	return taskCollection.Remove(bson.M{"name": name})
}

func UpdateTask(task *model.Task) error {
	return taskCollection.Update(bson.M{"name": task.Name}, task)
}

func AllTasks() ([]model.Task, error) {
	cnt, err := taskCollection.Count()
	if err != nil {
		return nil, err
	}

	allTasks := make([]model.Task, cnt)
	iter := taskCollection.Find(nil).Limit(1000).Iter()
	if err := iter.All(&allTasks); err != nil {
		return nil, err
	}
	return allTasks, nil
}

func ListTask(page int, limit int) ([]*model.Task, error) {
	return nil, nil
}

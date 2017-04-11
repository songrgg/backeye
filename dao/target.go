package dao

import (
	"github.com/songrgg/backeye/model"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func NewTarget(target *model.Target) error {
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

func GetTarget(name string) (*model.Target, error) {
	target := model.Target{}
	err := targetCollection.Find(bson.M{"name": name}).One(&target)
	return &target, err
}

func RemoveTarget(name string) error {
	return targetCollection.Remove(bson.M{"name": name})
}

func UpdateTarget(target *model.Target) error {
	return targetCollection.Update(bson.M{"name": target.Name}, target)
}

func All() ([]model.Target, error) {
	cnt, err := targetCollection.Count()
	if err != nil {
		return nil, err
	}

	allTargets := make([]model.Target, cnt)
	iter := targetCollection.Find(nil).Limit(1000).Iter()
	if err := iter.All(&allTargets); err != nil {
		return nil, err
	}
	return allTargets, nil
}

func ListTarget(page int, limit int) ([]*model.Target, error) {
	return nil, nil
}

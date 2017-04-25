package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/songrgg/backeye/model"
	"github.com/songrgg/backeye/model/form"
	"github.com/songrgg/backeye/std"
	modelmapper "gopkg.in/jeevatkm/go-model.v0"

	"time"
)

// NewTask2 creates a new task
func NewTask2(f *form.Task) error {
	nowTime := time.Now()
	task := model.Task{
		TimeMixin: std.TimeMixin{
			CreatedAt: nowTime,
			UpdatedAt: nowTime,
		},
	}
	errors := modelmapper.Copy(&task, f)
	if len(errors) > 0 {
		return errors[0]
	}

	// save task
	db := model.DB().Save(&task)
	if err := db.Error; err != nil {
		std.LogErrorc("mysql", err, "failed to create task")
		return err
	}

	if f.Watches != nil {
		for _, watch := range f.Watches {
			if err := NewWatch(task.ID, &watch); err != nil {
				std.LogErrorc("mysql", err, "failed to create watch")
				return err
			}
		}
	}
	return nil
}

// NewWatch creates new watch for specified task
func NewWatch(taskID int64, f *form.Watch) error {
	nowTime := time.Now()
	w := model.Watch{
		TaskID: taskID,
		TimeMixin: std.TimeMixin{
			CreatedAt: nowTime,
			UpdatedAt: nowTime,
		},
	}

	modelmapper.Copy(&w, f)
	db := model.DB().Save(&w)
	if err := db.Error; err != nil {
		std.LogErrorc("mysql", err, "failed to create watch")
		return err
	}

	for _, a := range f.Assertions {
		if err := NewAssertion(w.ID, &a); err != nil {
			std.LogErrorc("mysql", err, "failed to create assertion")
			return err
		}
	}
	return nil
}

// NewAssertion creates new assertion for specified watch
func NewAssertion(watchID int64, f *form.Assertion) error {
	nowTime := time.Now()
	a := model.Assertion{
		WatchID: watchID,
		TimeMixin: std.TimeMixin{
			CreatedAt: nowTime,
			UpdatedAt: nowTime,
		},
	}

	modelmapper.Copy(&a, f)
	db := model.DB().Save(&a)
	if err := db.Error; err != nil {
		return err
	}
	return nil
}

func RemoveTask(id int64) error {
	db := model.DB().Delete(nil, map[string]interface{}{"id": id})
	if err := db.Error; err != nil {
		return err
	}
	return nil
}

func UpdateTask(id int64, f *form.Task) error {
	task := model.Task{
		TimeMixin: std.TimeMixin{
			UpdatedAt: time.Now(),
		},
	}
	modelmapper.Copy(&task, f)

	db := model.DB().Find(&task, map[string]interface{}{"id": id}).Update(task)
	if err := db.Error; err != nil {
		return err
	}
	return nil
}

func GetTask(id int64) (*model.Task, error) {
	var task model.Task
	db := model.DB().Where(map[string]interface{}{"id": id}).First(&task)
	if err := db.Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func AllTasks() ([]model.Task, error) {
	tasks := make([]model.Task, 0)
	db := model.DB().Find(&tasks)
	if err := db.Error; err != nil {
		return nil, err
	}

	for i := range tasks {
		assembleTask(db, &tasks[i])
	}
	return tasks, nil
}

func assembleTask(db *gorm.DB, t *model.Task) error {
	var watches []model.Watch
	db.Model(t).Related(&watches)

	for i := range watches {
		assembleWatch(db, &watches[i])
	}

	t.Watches = watches
	return nil
}

func assembleWatch(db *gorm.DB, w *model.Watch) error {
	var assertions []model.Assertion
	db.Model(w).Related(&assertions)
	w.Assertions = assertions
	return nil
}

func ListTask(maxID string, limit int) ([]model.Task, error) {
	tasks := make([]model.Task, 0)
	db := model.DB().Find(&tasks)
	if err := db.Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

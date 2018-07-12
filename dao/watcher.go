package dao

import (
	"time"

	"encoding/json"

	"github.com/songrgg/backeye/model"
	"github.com/songrgg/backeye/model/form"
	"github.com/songrgg/backeye/std"
	modelmapper "gopkg.in/jeevatkm/go-model.v0"
)

// NewWatcher creates a new watcher
func NewWatcher(f *form.Watcher) error {
	nowTime := time.Now()
	watcher := model.Watcher{
		TimeMixin: std.TimeMixin{
			CreatedAt: nowTime,
			UpdatedAt: nowTime,
		},
	}
	errors := modelmapper.Copy(&watcher, f)

	// TODO some validation

	pointConf, _ := json.Marshal(f.Points)
	watcher.Points = string(pointConf)
	if len(errors) > 0 {
		return errors[0]
	}

	// save task
	model.DB().Save(&watcher)
	if err := model.DB().Error; err != nil {
		std.LogErrorc("mysql", err, "failed to create project")
		return err
	}
	return nil
}

func GetWatchers() ([]model.Watcher, error) {
	var watchers []model.Watcher
	model.DB().Find(&watchers, "disabled = false").Order("updated_at desc")
	return watchers, nil
}

func GetWatcher(ID int64) (model.Watcher, error) {
	var watcher model.Watcher
	model.DB().First(&watcher, "id = ? and disabled = false", ID).Order("updated_at desc")
	return watcher, nil
}

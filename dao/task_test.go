package dao

import (
	"testing"

	"github.com/songrgg/backeye/model"
	"github.com/songrgg/backeye/model/form"
	"github.com/stretchr/testify/assert"
	modelmapper "gopkg.in/jeevatkm/go-model.v0"
)

func TestModelMapping(t *testing.T) {
	f := form.Task{
		ProjectID: 1,
		Name:      "apitest",
		Type:      "cron",
		Status:    "paused",
		Desc:      "apitest is cool",
		CronSpec:  "*/2 * * * *",
	}
	task := model.Task{}
	modelmapper.Copy(&task, f)
	assert.Equal(t, task.ProjectID, f.ProjectID)
	assert.Equal(t, task.Name, f.Name)
	assert.Equal(t, task.Type, f.Type)
	assert.Equal(t, task.Status, f.Status)
	assert.Equal(t, task.Desc, f.Desc)
	assert.Equal(t, task.CronSpec, f.CronSpec)
}

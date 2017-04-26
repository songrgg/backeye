package parser

import (
	"testing"

	"github.com/songrgg/backeye/model"
	"github.com/stretchr/testify/assert"
)

func TestRunSuccess(t *testing.T) {
	task := model.Task{
		ID:        1,
		ProjectID: 1,
		Name:      "apitest",
		Type:      "normal",
		Status:    "active",
		Desc:      "api test demo",
		CronSpec:  "*/2 * * * *",
		Watches: []model.Watch{
			model.Watch{
				ID:       1,
				TaskID:   1,
				Name:     "post API",
				Desc:     "fetch latest posts",
				Interval: 0,
				Path:     "https://api-prod.wallstreetcn.com/apiv1/content/articles",
				Method:   "GET",
				Headers: `{
					"User-Agent": "backeye"
				}`,
				Assertions: []model.Assertion{
					model.Assertion{
						ID:       1,
						WatchID:  1,
						Type:     "header",
						Code:     "",
						Source:   "header",
						Operator: "equal",
						Left:     "status_code",
						Right:    "200",
						Revision: 100,
					},
				},
			},
		},
	}

	defaultParser := DefaultParser{}
	assert.NotNil(t, defaultParser)

	newTask, err := defaultParser.Translate(&task)
	assert.Nil(t, err)
	assert.Equal(t, task.ID, newTask.ID)
	assert.Equal(t, task.CronSpec, newTask.CronSpec)
	assert.Equal(t, task.Name, newTask.Name)
	assert.Equal(t, task.Desc, newTask.Desc)

	assert.Len(t, newTask.Watches, 1)
}

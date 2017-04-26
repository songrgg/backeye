package parser

import (
	"testing"

	"github.com/songrgg/backeye/model"
	"github.com/songrgg/backeye/watch/http"
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

func TestNewVariable(t *testing.T) {
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
				Variables: []model.Variable{
					model.Variable{
						ID:      1,
						WatchID: 1,
						Name:    "postID",
						Value:   "$RESPONSE.data.items[0].id",
					},
				},
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
			model.Watch{
				ID:       2,
				TaskID:   1,
				Name:     "post detail",
				Desc:     "fetch specified post",
				Interval: 0,
				Path:     "https://api-prod.wallstreetcn.com/apiv1/content/articles/${postID}?extract=0",
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
						Revision: 1,
					},
				},
			},
		},
	}

	defaultParser := DefaultParser{}
	newTask, err := defaultParser.ParseTask(&task)
	assert.NoError(t, err)
	assert.Len(t, newTask.Watches, 2)

	httpWatch, ok := newTask.Watches[0].(*http.Watch)
	assert.True(t, ok)
	assert.NotNil(t, httpWatch)
	assert.Len(t, httpWatch.Variables, 1)
	assert.Equal(t, "postID", httpWatch.Variables[0].Name)
	assert.Equal(t, "$RESPONSE.data.items[0].id", httpWatch.Variables[0].Value)

	// vm := otto.New()
	// vm.Set("task", map[string]interface{}{
	// 	"name": "apitest",
	// })
	// val, _ := vm.Run(`task.name`)
	// fmt.Println(val)
}

package json

import (
	"context"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	jsonParser := Parser{}
	assert.Nil(t, jsonParser.Task)

	bytes, err := ioutil.ReadFile("success.json")
	assert.NoError(t, err)

	jsonParser.load(bytes)
	assert.Equal(t, "Post API", jsonParser.Task.Name)
	assert.Equal(t, "post API", jsonParser.Task.Desc)
	assert.Equal(t, "*/2 * * * *", jsonParser.Task.Cron)
	assert.Equal(t, 1, len(jsonParser.Task.Watches))
	assert.Equal(t, "post list", jsonParser.Task.Watches[0].Name)
	assert.Equal(t, "post list", jsonParser.Task.Watches[0].Desc)
	assert.Equal(t, 0, jsonParser.Task.Watches[0].Interval)
	assert.Equal(t, "https://api-prod.wallstreetcn.com/apiv1/content/articles", jsonParser.Task.Watches[0].Path)
	assert.Equal(t, "GET", jsonParser.Task.Watches[0].Method)
	assert.Len(t, jsonParser.Task.Watches[0].Headers, 1)
	assert.Len(t, jsonParser.Task.Watches[0].Assertions, 2)
}

func TestTranslate(t *testing.T) {
	jsonParser := Parser{}
	assert.Nil(t, jsonParser.Task)

	bytes, err := ioutil.ReadFile("success.json")
	assert.NoError(t, err)

	jsonParser.load(bytes)
	task := jsonParser.Task
	assert.Equal(t, "Post API", task.Name)
	assert.Equal(t, "post API", task.Desc)
}

func TestRunSuccess(t *testing.T) {
	jsonParser := Parser{}
	assert.Nil(t, jsonParser.Task)

	bytes, err := ioutil.ReadFile("success.json")
	assert.NoError(t, err)

	task, err := jsonParser.Translate(bytes)
	assert.Equal(t, "Post API", task.Name)
	assert.Equal(t, "post API", task.Desc)
	assert.Len(t, task.Watches, 1)
	assert.NoError(t, err)

	watchResults, err := task.Run(context.Background())
	assert.Nil(t, err)
	assert.Len(t, watchResults[0].Assertions, 2)
	assert.True(t, watchResults[0].Assertions[0].Passed)
	assert.True(t, watchResults[0].Assertions[1].Passed)
}

func TestRunFailure(t *testing.T) {
	jsonParser := Parser{}
	assert.Nil(t, jsonParser.Task)

	bytes, err := ioutil.ReadFile("failure.json")
	assert.NoError(t, err)

	task, err := jsonParser.Translate(bytes)
	assert.Equal(t, "Post API", task.Name)
	assert.Equal(t, "post API", task.Desc)
	assert.Len(t, task.Watches, 1)
	assert.NoError(t, err)

	watchResults, err := task.Run(context.Background())
	assert.Nil(t, err)
	assert.Len(t, watchResults, 1)
	assert.Len(t, watchResults[0].Assertions, 2)
	assert.True(t, watchResults[0].Assertions[0].Passed)
	assert.False(t, watchResults[0].Assertions[1].Passed)
}

func TestRunMultipleAssertions(t *testing.T) {
	jsonParser := Parser{}
	assert.Nil(t, jsonParser.Task)

	bytes, err := ioutil.ReadFile("multiple_assertions.json")
	assert.NoError(t, err)

	task, err := jsonParser.Translate(bytes)
	assert.NoError(t, err)
	assert.Equal(t, "Post API", task.Name)
	assert.Equal(t, "post API", task.Desc)
	assert.Len(t, task.Watches, 1)

	watchResults, err := task.Run(context.Background())
	assert.Nil(t, err)
	assert.Len(t, watchResults, 1)
	assert.Len(t, watchResults[0].Assertions, 3)
	assert.True(t, watchResults[0].Assertions[0].Passed)
	assert.True(t, watchResults[0].Assertions[1].Passed)
	assert.True(t, watchResults[0].Assertions[2].Passed)
}

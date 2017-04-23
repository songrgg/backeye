package schedule

import (
	"errors"

	"github.com/labstack/echo"
	"github.com/songrgg/backeye/dao"
	"github.com/songrgg/backeye/helper"
	"github.com/songrgg/backeye/model"
)

func HTTPAddTask(ctx echo.Context) error {
	args := &model.Task{}
	if err := ctx.Bind(args); err != nil {
		return helper.ErrorResponse(ctx, errors.New("fail to bind args"))
	}

	return nil
}

func HTTPGetTasks(ctx echo.Context) error {
	task := listTasks(ctx)
	return helper.SuccessResponse(ctx, task)
}

func HTTPGetTask(ctx echo.Context) error {
	id := ctx.Param("id")
	if id == "" {
		return ctx.JSON(200, helper.Payload(nil))
	}
	task := getTask(ctx, id)
	return helper.SuccessResponse(ctx, task)
}

func HTTPDeleteTask(ctx echo.Context) error {
	id := ctx.Param("id")
	dao.RemoveTask(id)
	return helper.SuccessResponse(ctx, helper.Payload(nil))
}

func HTTPUpdateTask(ctx echo.Context) error {
	return nil
}

func HTTPGetWatchResults(ctx echo.Context) error {
	id := ctx.Param("id")
	watchResults := getWatchResults(ctx, id, "")
	return helper.SuccessResponse(ctx, watchResults)
}

func HTTPGetTaskHealth(ctx echo.Context) error {
	id := ctx.Param("id")
	return helper.SuccessResponse(ctx, getTaskHealth(ctx, id))
}

func listTasks(ctx echo.Context) *model.Task {
	task, err := dao.GetTask("Post API")
	if err != nil {
		return nil
	}
	return task
}

func getTask(ctx echo.Context, id string) *model.Task {
	task, err := dao.GetTask(id)
	if err != nil {
		return nil
	}
	return task
}

func getWatchResults(ctx echo.Context, taskName string, watchName string) []model.WatchResult {
	results, err := dao.AllWatchResults(taskName)
	if err != nil {
		return nil
	}
	return results
}

func getTaskHealth(ctx echo.Context, taskName string) *model.TaskHealth {
	results, err := dao.AllWatchResults(taskName)
	if err != nil {
		return nil
	}

	total := 0
	success := 0
	for _, result := range results {
		for _, assertion := range result.Assertions {
			if assertion.Success {
				success++
			}
			total++
		}
	}
	return &model.TaskHealth{
		Total:   total,
		Success: success,
	}
}

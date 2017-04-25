package schedule

import (
	"github.com/labstack/echo"
	"github.com/songrgg/backeye/dao"
	"github.com/songrgg/backeye/helper"
	"github.com/songrgg/backeye/model"
	"github.com/songrgg/backeye/model/form"
	"github.com/songrgg/backeye/std"
)

func HTTPAddTask(ctx echo.Context) error {
	args := &form.Task{}
	if err := ctx.Bind(args); err != nil {
		std.LogErrorc("echo", err, "failed to bind args")
		return helper.ErrorResponse(ctx, err)
	}

	err := dao.NewTask2(args)
	if err != nil {
		std.LogErrorc("backeye", err, "failed to create task")
		return helper.ErrorResponse(ctx, err)
	}

	return helper.SuccessResponse(ctx, nil)
}

func HTTPGetTasks(ctx echo.Context) error {
	tasks, err := listTasks(ctx)
	if err != nil {
		return helper.ErrorResponse(ctx, err)
	}
	return helper.SuccessResponse(ctx, tasks)
}

func HTTPGetTask(ctx echo.Context) error {
	id := std.GetID(ctx)
	task, err := getTask(ctx, id)
	if err != nil {
		return helper.ErrorResponse(ctx, err)
	}
	return helper.SuccessResponse(ctx, task)
}

func HTTPDeleteTask(ctx echo.Context) error {
	id := std.GetID(ctx)
	err := dao.RemoveTask(id)
	if err != nil {
		return helper.ErrorResponse(ctx, err)
	}
	return helper.SuccessResponse(ctx, nil)
}

func HTTPUpdateTask(ctx echo.Context) error {
	id := std.GetID(ctx)
	args := &form.Task{}
	if err := ctx.Bind(args); err != nil {
		std.LogErrorc("echo", err, "failed to bind args")
		return helper.ErrorResponse(ctx, err)
	}

	err := dao.UpdateTask(id, args)
	if err != nil {
		return helper.ErrorResponse(ctx, err)
	}
	return helper.SuccessResponse(ctx, nil)
}

func HTTPGetWatchResults(ctx echo.Context) error {
	id := ctx.Param("id")
	watchResults := getWatchResults(ctx, id, "")
	return helper.SuccessResponse(ctx, watchResults)
}

func HTTPGetTaskHealth(ctx echo.Context) error {
	return nil
}

func listTasks(ctx echo.Context) ([]model.Task, error) {
	maxID := std.FetchStrParam(ctx, "maxID", "")
	limit := std.FetchIntParam(ctx, "limit", 10)
	tasks, err := dao.ListTask(maxID, limit)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func getTask(ctx echo.Context, id int64) (*model.Task, error) {
	return dao.GetTask(id)
}

func getWatchResults(ctx echo.Context, taskName string, watchName string) []model.WatchResult {
	maxID := std.FetchStrParam(ctx, "maxID", "")
	limit := std.FetchIntParam(ctx, "limit", 10)
	results, err := dao.AllWatchResults(taskName, maxID, limit)
	if err != nil {
		return nil
	}
	return results
}

// func getTaskHealth(ctx echo.Context, taskName string) *model.TaskHealth {
// 	maxID := std.FetchStrParam(ctx, "maxID", "")
// 	limit := std.FetchIntParam(ctx, "limit", 10)
// 	results, err := dao.AllWatchResults(taskName, maxID, limit)
// 	if err != nil {
// 		return nil
// 	}

// 	total := 0
// 	success := 0
// 	for _, result := range results {
// 		for _, assertion := range result.Assertions {
// 			if assertion.Success {
// 				success++
// 			}
// 			total++
// 		}
// 	}
// 	return &model.TaskHealth{
// 		Total:   total,
// 		Success: success,
// 	}
// }

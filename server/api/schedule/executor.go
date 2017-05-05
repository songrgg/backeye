package schedule

import (
	"strconv"

	"github.com/labstack/echo"
	"github.com/songrgg/backeye/dao"
	"github.com/songrgg/backeye/helper"
	"github.com/songrgg/backeye/model/form"
	"github.com/songrgg/backeye/schedule"
	"github.com/songrgg/backeye/std"
)

// @Title executor
// @Description run task for API testing
// @Param   id     			path    	int   	  	true 		"task id"
// @Accept  json
// @Resource executor
// @Router /v1/executors/runtask/{id} [post]
func HTTPRunTaskByID(ctx echo.Context) error {
	id := ctx.Param("id")
	ID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		std.LogErrorc("backeye", err, "failed to parse ID")
		return helper.ErrorResponse(ctx, err)
	}

	// execute the task once
	taskModel, err := dao.GetTask(ID)
	if err != nil {
		std.LogErrorc("echo", err, "failed to run task")
		return helper.ErrorResponse(ctx, err)
	}

	results, err := schedule.RunTask(taskModel)
	if err != nil {
		std.LogErrorc("echo", err, "failed to run task")
		return helper.ErrorResponse(ctx, err)
	}

	return helper.SuccessResponse(ctx, results)
}

// @Title executor
// @Description run task for API testing
// @Param   args	body  form.Task    true	"API testing task"
// @Accept  json
// @Resource executor
// @Router /v1/executors/runtask [post]
func HTTPRunTask(ctx echo.Context) error {
	args := &form.Task{}
	if err := ctx.Bind(args); err != nil {
		std.LogErrorc("echo", err, "failed to bind args")
		return helper.ErrorResponse(ctx, err)
	}

	// execute the task once
	taskModel, err := form.ParseTask(args)
	if err != nil {
		std.LogErrorc("echo", err, "failed to parse task")
	}

	results, err := schedule.RunTask(&taskModel)
	if err != nil {
		std.LogErrorc("echo", err, "failed to run task")
		return helper.ErrorResponse(ctx, err)
	}

	return helper.SuccessResponse(ctx, results)
}

// @Title executor
// @Description run task for API testing
// @Param   args	body  form.Task    true	"API testing task"
// @Accept  json
// @Resource executor
// @Router /v1/executors/verifytask [post]
func HTTPVerifyTask(ctx echo.Context) error {
	args := &form.Task{}
	if err := ctx.Bind(args); err != nil {
		std.LogErrorc("echo", err, "failed to bind args")
		return helper.ErrorResponse(ctx, err)
	}

	// TODO execute the task once

	return helper.SuccessResponse(ctx, nil)
}

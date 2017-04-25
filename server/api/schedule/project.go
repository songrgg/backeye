package schedule

import (
	"strconv"

	"github.com/labstack/echo"
	"github.com/songrgg/backeye/dao"
	"github.com/songrgg/backeye/helper"
	"github.com/songrgg/backeye/model/form"
	"github.com/songrgg/backeye/std"
)

func httpAddProject(ctx echo.Context) error {
	args := &form.Project{}
	if err := ctx.Bind(args); err != nil {
		std.LogErrorc("echo", err, "failed to bind args")
		return helper.ErrorResponse(ctx, err)
	}

	err := dao.NewProject(args)
	if err != nil {
		std.LogErrorc("backeye", err, "failed to create project")
		return helper.ErrorResponse(ctx, err)
	}

	return helper.SuccessResponse(ctx, nil)
}

func httpGetProjects(ctx echo.Context) error {
	projs, err := dao.GetProjects()
	if err != nil {
		std.LogErrorc("backeye", err, "failed to fetch project")
		return helper.ErrorResponse(ctx, err)
	}

	return helper.SuccessResponse(ctx, projs)
}

func httpGetProject(ctx echo.Context) error {
	id := ctx.Param("id")
	ID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		std.LogErrorc("backeye", err, "failed to parse ID")
		return helper.ErrorResponse(ctx, err)
	}

	proj, err := dao.GetProject(ID)
	if err != nil {
		std.LogErrorc("backeye", err, "failed to fetch project")
		return helper.ErrorResponse(ctx, err)
	}
	return helper.SuccessResponse(ctx, proj)
}

func httpDeleteProject(ctx echo.Context) error {
	id := ctx.Param("id")
	ID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		std.LogErrorc("backeye", err, "failed to parse ID")
		return helper.ErrorResponse(ctx, err)
	}

	err = dao.RemoveProject(ID)
	if err != nil {
		std.LogErrorc("backeye", err, "failed to fetch project")
		return helper.ErrorResponse(ctx, err)
	}

	return helper.SuccessResponse(ctx, nil)
}

func httpUpdateProject(ctx echo.Context) error {
	id := ctx.Param("id")
	ID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		std.LogErrorc("backeye", err, "failed to parse ID")
		return helper.ErrorResponse(ctx, err)
	}

	args := &form.Project{}
	if err := ctx.Bind(args); err != nil {
		std.LogErrorc("echo", err, "failed to bind args")
		return helper.ErrorResponse(ctx, err)
	}

	err = dao.UpdateProject(ID, args)
	if err != nil {
		std.LogErrorc("backeye", err, "failed to fetch project")
		return helper.ErrorResponse(ctx, err)
	}
	return helper.SuccessResponse(ctx, nil)
}

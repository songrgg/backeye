package schedule

import (
	"errors"

	"github.com/labstack/echo"
	"github.com/songrgg/backeye/dao"
	"github.com/songrgg/backeye/helper"
	"github.com/songrgg/backeye/model"
)

func HTTPAddSchedule(ctx echo.Context) error {
	args := &model.Target{}
	if err := ctx.Bind(args); err != nil {
		return helper.ErrorResponse(ctx, errors.New("fail to bind args"))
	}

	return nil
}

func HTTPGetSchedules(ctx echo.Context) error {
	target := listSchedules(ctx)
	return helper.SuccessResponse(ctx, target)
}

func HTTPGetSchedule(ctx echo.Context) error {
	id := ctx.Param("id")
	if id == "" {
		return ctx.JSON(200, helper.Payload(nil))
	}
	target := getSchedule(ctx, id)
	return helper.SuccessResponse(ctx, target)
}

func HTTPDeleteSchedule(ctx echo.Context) error {
	id := ctx.Param("id")
	dao.RemoveTarget(id)
	return helper.SuccessResponse(ctx, helper.Payload(nil))
}

func HTTPUpdateSchedule(ctx echo.Context) error {
	return nil
}

func listSchedules(ctx echo.Context) *model.Target {
	target, err := dao.GetTarget("Post API")
	if err != nil {
		return nil
	}
	return target
}

func getSchedule(ctx echo.Context, id string) *model.Target {
	target, err := dao.GetTarget(id)
	if err != nil {
		return nil
	}
	return target
}

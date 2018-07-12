package watcher

import (
	"strconv"

	"github.com/labstack/echo"
	"github.com/songrgg/backeye/dao"
	"github.com/songrgg/backeye/helper"
)

// @Title Watcher list
// @Description fetch watcher list
// @Accept  json
// @Resource watcher
// @Router /v1/watchers [get]
func HTTPGetWatchers(ctx echo.Context) error {
	watchers, err := dao.GetWatchers()
	if err != nil {
		return helper.ErrorResponse(ctx, err)
	}

	return helper.SuccessResponse(ctx, watchers)
}

// @Title Watcher detail
// @Description fetch watcher detail
// @Accept  json
// @Resource watcher
// @Router /v1/watchers/{id} [get]
func HTTPGetWatcher(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return helper.ErrorResponse(ctx, err)
	}

	watchers, err := dao.GetWatcher(id)
	if err != nil {
		return helper.ErrorResponse(ctx, err)
	}

	return helper.SuccessResponse(ctx, watchers)
}

// @Title Watcher detail
// @Description fetch watcher detail
// @Accept  json
// @Resource watcher
// @Router /v1/watchers/{id}/assertions [get]
func HTTPGetWatcherAssertions(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return helper.ErrorResponse(ctx, err)
	}

	results, err := dao.GetAssertionResults(id, "created_at desc", 20)
	if err != nil {
		return helper.ErrorResponse(ctx, err)
	}

	return helper.SuccessResponse(ctx, results)
}

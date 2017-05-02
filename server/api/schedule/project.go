package schedule

import (
	"strconv"

	"github.com/labstack/echo"
	"github.com/songrgg/backeye/dao"
	"github.com/songrgg/backeye/helper"
	"github.com/songrgg/backeye/model/form"
	"github.com/songrgg/backeye/std"
)

// @Title project
// @Description add project
// @Param   args	body  form.Project    true	"API testing project"
// @Accept  json
// @Resource project
// @Router /v1/projects [post]
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

// @Title project
// @Description fetch project list
// @Accept  json
// @Resource project
// @Router /v1/projects [get]
func httpGetProjects(ctx echo.Context) error {
	projs, err := dao.GetProjects()
	if err != nil {
		std.LogErrorc("backeye", err, "failed to fetch project")
		return helper.ErrorResponse(ctx, err)
	}

	return helper.SuccessResponse(ctx, projs)
}

// @Title project
// @Description fetch the specified project
// @Param   id     			path    	int   	  	true 		"project id"
// @Accept  json
// @Resource project
// @Router /v1/projects/{id} [get]
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

// @Title project
// @Description delete the specified project
// @Param   id     			path    	int   	  	true 		"project id"
// @Accept  json
// @Resource project
// @Router /v1/projects/{id} [delete]
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

// @Title project
// @Description update the specified project
// @Param   id     			path    	int   	  	true 		"project id"
// @Param   args	body  form.Project    true	"API testing project"
// @Accept  json
// @Resource project
// @Router /v1/projects/{id} [put]
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

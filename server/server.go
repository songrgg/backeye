package server

import (
	"github.com/labstack/echo"
	"github.com/songrgg/backeye/common"
	"github.com/songrgg/backeye/server/api/schedule"
	"github.com/songrgg/backeye/server/middleware"
)

var app = echo.New()

func init() {
	app.Use(middleware.RequestCORS())

	mountAPIModule(app)
}

// RunServer starts a server
func RunServer() {
	if common.Config.CertPem != "" && common.Config.KeyPem != "" {
		app.StartTLS(common.Config.Bind, common.Config.CertPem, common.Config.KeyPem)
	} else {
		app.Start(common.Config.Bind)
	}
}

func mountAPIModule(e *echo.Echo) {
	apiv1 := e.Group("/v1")
	schedule.MountAPI(apiv1)
}

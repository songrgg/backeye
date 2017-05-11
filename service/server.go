package server

import (
	"github.com/labstack/echo"
	"github.com/songrgg/backeye/service/api/schedule"
	"github.com/songrgg/backeye/service/middleware"
	"github.com/songrgg/backeye/service/swagger"
	"github.com/songrgg/backeye/std"
)

var app = echo.New()

func init() {
	app.Use(middleware.RequestCORS())

	mountAPIModule(app)
}

// RunServer starts a server
func RunServer() {
	if std.Config.CertPem != "" && std.Config.KeyPem != "" {
		app.StartTLS(std.Config.Bind, std.Config.CertPem, std.Config.KeyPem)
	} else {
		app.Start(std.Config.Bind)
	}
}

func mountAPIModule(e *echo.Echo) {
	apiv1 := e.Group("/v1")
	schedule.MountAPI(apiv1)
	swagger.MountSwaggerAPI(e.Group(""))
}

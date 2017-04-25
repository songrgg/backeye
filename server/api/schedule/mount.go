package schedule

import (
	"github.com/labstack/echo"
)

// MountAPI registers API
func MountAPI(group *echo.Group) {
	mountTaskAPI(group)
	moutProjectAPI(group)
}

func mountTaskAPI(group *echo.Group) {
	task := group.Group("/tasks")
	task.POST("", HTTPAddTask)
	task.GET("", HTTPGetTasks)
	task.DELETE("/:id", HTTPDeleteTask)
	task.PUT("/:id", HTTPUpdateTask)
	task.GET("/:id", HTTPGetTask)
	task.GET("/:id/watch_results", HTTPGetWatchResults)
	task.GET("/:id/health", HTTPGetTaskHealth)
}

func moutProjectAPI(group *echo.Group) {
	proj := group.Group("/projects")
	proj.POST("", httpAddProject)
	proj.GET("", httpGetProjects)
	proj.GET("/:id", httpGetProject)
	proj.DELETE("/:id", httpDeleteProject)
	proj.PUT("/:id", httpUpdateProject)
}

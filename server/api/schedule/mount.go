package schedule

import (
	"github.com/labstack/echo"
)

// MountAPI registers API
func MountAPI(group *echo.Group) {
	mountTaskAPI(group)
}

func mountTaskAPI(group *echo.Group) {
	task := group.Group("/tasks")
	task.POST("/", HTTPAddTask)
	task.GET("/", HTTPGetTasks)
	task.DELETE("/", HTTPDeleteTask)
	task.PUT("/", HTTPUpdateTask)
	task.GET("/:id/watch_results", HTTPGetWatchResults)
	task.GET("/:id/health", HTTPGetTaskHealth)
}

package watcher

import (
	"github.com/labstack/echo"
)

// MountAPI registers API
func MountAPI(group *echo.Group) {
	mountWatcherAPI(group)
}

func mountWatcherAPI(group *echo.Group) {
	task := group.Group("/watchers")
	task.GET("", HTTPGetWatchers)
	task.GET("/:id", HTTPGetWatcher)
	task.GET("/:id/assertions", HTTPGetWatcherAssertions)
}

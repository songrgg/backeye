package schedule

import (
	"github.com/labstack/echo"
)

// MountAPI registers API
func MountAPI(group *echo.Group) {
	schedule := group.Group("/schedules")
	schedule.POST("/", HTTPAddSchedule)
	schedule.GET("/test", HTTPGetSchedules)
	schedule.DELETE("/", HTTPDeleteSchedule)
	schedule.PUT("/", HTTPUpdateSchedule)
}

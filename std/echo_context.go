package std

import (
	"strconv"

	"github.com/labstack/echo"
	"gitlab.wallstcn.com/wscnbackend/ivankastd"
)

func FetchStrParam(ctx echo.Context, name string, defaultVal string) string {
	val := ctx.QueryParam(name)
	if val == "" && defaultVal != "" {
		return defaultVal
	}
	return val
}

func FetchIntParam(ctx echo.Context, name string, defaultVal int) int {
	val := ctx.QueryParam(name)
	if val == "" {
		return defaultVal
	}

	i, err := strconv.Atoi(val)
	if err != nil {
		ivankastd.LogErrorc("strconv", err, "failed to parse int")
		return 0
	}

	return i
}

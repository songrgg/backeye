package std

import (
	"strconv"

	"github.com/labstack/echo"
)

func FetchStrParam(ctx echo.Context, name string, defaultVal string) string {
	val := ctx.QueryParam(name)
	if val == "" && defaultVal != "" {
		return defaultVal
	}
	return val
}

func GetID(ctx echo.Context) int64 {
	id := ctx.Param("id")
	ID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		LogErrorc("backeye", err, "failed to parse ID")
		return 0
	}
	return ID
}

func FetchIntParam(ctx echo.Context, name string, defaultVal int) int {
	val := ctx.QueryParam(name)
	if val == "" {
		return defaultVal
	}

	i, err := strconv.Atoi(val)
	if err != nil {
		LogErrorc("strconv", err, "failed to parse int")
		return 0
	}

	return i
}

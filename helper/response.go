package helper

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/songrgg/backeye/std"
)

// Response API body
type Response struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

func SuccessResponse(ctx echo.Context, data interface{}) error {
	var buffer bytes.Buffer
	m := json.NewEncoder(&buffer)
	if err := m.Encode(data); err != nil {
		return ErrorResponse(ctx, errors.New("Illegal JSON"))
	}
	return ctx.JSON(http.StatusOK, Payload(buffer.Bytes()))
}

func ErrorResponse(ctx echo.Context, err error) error {
	if err != nil {
		return ctx.JSON(http.StatusOK, Payload([]byte("{}"), "50000", err.Error()))
	}
	return ctx.JSON(http.StatusOK, Payload([]byte("{}"), "50000", ""))
}

func Payload(data json.RawMessage, fields ...string) *Response {
	res := &Response{
		Code:    std.ErrOK.Code,
		Message: std.ErrOK.Message,
		Data:    data,
	}

	length := len(fields)
	if length >= 1 {
		singleCode, _ := strconv.Atoi(fields[0])
		res.Code = int(singleCode)
	}
	if length >= 2 {
		res.Message = fields[1]
	}

	return res
}

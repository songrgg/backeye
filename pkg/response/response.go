package response

import (
	"context"

	"github.com/songrgg/backeye/pkg/common"
)

type Response map[string]interface{}

func SetResponse(ctx context.Context, res Response) context.Context {
	context.WithValue(ctx, common.ResponseKey, res)
	return nil
}

func GetResponse(ctx context.Context) Response {
	return ctx.Value(common.ResponseKey).(Response)
}

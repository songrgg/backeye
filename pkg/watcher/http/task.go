package http

import (
	"context"
	"io/ioutil"
	nethttp "net/http"
	"time"

	"errors"

	"fmt"
	"strings"

	"github.com/songrgg/backeye/pkg/common"
	"github.com/songrgg/backeye/pkg/response"
)

type Task struct {
	Path    string        `json:"path"`
	Method  string        `json:"method"`
	Timeout time.Duration `json:"timeout"`
	//Vars    map[string]string
}

const (
	Get  = "get"
	Post = "post"
	Put  = "put"
	Head = "head"
)

func (h *Task) Do(ctx context.Context) (response.Response, error) {
	var (
		res       = make(response.Response)
		startTime = time.Now()
		resp      *nethttp.Response
		err       error
		body      []byte
		latency   time.Duration
		path      = h.Path
	)

	vars := ctx.Value(common.VariablesKey)
	if vars != nil {
		varsMap := vars.(map[string]interface{})
		for k, v := range varsMap {
			replaceStr := fmt.Sprintf("{{%s}}", k)
			for strings.Contains(path, replaceStr) {
				path = strings.Replace(path, replaceStr, "%v", 1)
				path = fmt.Sprintf(path, v)
			}
		}
	}

	switch h.Method {
	case Get:
		resp, err = nethttp.Get(path)
		if err != nil {
			return nil, err
		}

		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

	case Head:
		resp, err = nethttp.Head(path)
		if err != nil {
			return nil, err
		}

		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

	default:
		return nil, errors.New("method not implemented")
	}

	latency = time.Since(startTime)

	// Set headers, latency, body
	SetStatus(res, resp.StatusCode)
	SetHeaders(res, resp.Header)
	SetLatency(res, latency)
	SetBody(res, body)

	return res, nil
}

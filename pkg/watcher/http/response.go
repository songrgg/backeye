package http

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/songrgg/backeye/pkg/response"
)

func SetStatus(res response.Response, status int) {
	res["status"] = status
}

func GetStatus(res response.Response) int {
	v, err := get(res, "status")
	if err != nil {
		return 0
	}
	return v.(int)
}

func SetHeaders(res response.Response, headers http.Header) {
	res["headers"] = headers
}

func GetHeader(res response.Response, key string) string {
	if v := GetHeaders(res); v != nil {
		return v.Get(key)
	}
	return ""
}

func GetHeaders(res response.Response) http.Header {
	v, err := get(res, "headers")
	if err != nil {
		return nil
	}
	return v.(http.Header)
}

func SetLatency(res response.Response, latency time.Duration) {
	res["latency"] = strconv.FormatInt(latency.Nanoseconds(), 10)
}

func GetLatency(res response.Response) time.Duration {
	v, err := get(res, "latency")
	if err != nil {
		return 0
	}

	l, _ := strconv.Atoi(v.(string))
	return time.Duration(l) * time.Nanosecond
}

func SetBody(res response.Response, body []byte) {
	res["body"] = body
}

func GetBody(res response.Response) []byte {
	v, err := get(res, "body")
	if err != nil {
		return nil
	}
	return v.([]byte)
}

func get(res response.Response, key string) (interface{}, error) {
	if v, ok := res[key]; !ok {
		return nil, errors.New("key not exist")
	} else {
		return v, nil
	}
}

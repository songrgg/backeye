package conf

import "github.com/songrgg/backeye/pkg/watcher/http"

type Config struct {
	Task       http.Task            `json:"task"`
	Assertions []http.AssertionConf `json:"assertions"`
	Variables  []http.Variable      `json:"variables"`
}

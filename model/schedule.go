package model

type Task struct {
	Name     string  `json:"name"`
	Desc     string  `json:"desc"`
	CronSpec string  `json:"cron"`
	Watches  []Watch `json:"watches"`
	Enabled  bool    `json:"enabled"`
}

type Watch struct {
	Name       string            `json:"name"`
	Desc       string            `json:"desc"`
	Interval   int32             `json:"interval"`
	Path       string            `json:"path"`
	Method     string            `json:"method"`
	Headers    map[string]string `json:"headers"`
	Assertions []Assertion       `json:"assertions"`
}

type Assertion struct {
	Type     string `json:"type"`
	Code     string `json:"code"`
	Timeout  string `json:"timeout"` // TODO: add timeout support
	Source   string `json:"source"`
	Operator string `json:"operator"`
	Left     string `json:"left"`
	Right    string `json:"right"`
}

// AssertionResult indicates the assertion result
type AssertionResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// WatchResult indicates the watch's result
type WatchResult struct {
	TaskName      string            `json:"task_name"`
	WatchName     string            `json:"watch_name"`
	ExecutionTime int64             `json:"execution_time"`
	Assertions    []AssertionResult `json:"assertion_results"`
}

// TaskHealth indicates the task's health condition
type TaskHealth struct {
	Total   int `json:"total"`
	Success int `json:"success"`
}

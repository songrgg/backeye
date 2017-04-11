package model

type Target struct {
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

type AssertionResult struct {
	AssertionID int64
	Success     bool
	Message     string
}

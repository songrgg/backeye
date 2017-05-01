package form

type Project struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type Task struct {
	ProjectID int64      `json:"project_id"`
	Name      string     `json:"name"`
	Type      string     `json:"type"`
	Status    string     `json:"status"`
	Desc      string     `json:"desc"`
	CronSpec  string     `json:"cron"`
	Watches   []Watch    `json:"watches"`
	Variables []Variable `json:"variables"`
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

type Variable struct {
	Name  string `json:"name"`
	Value string `json:"value"`
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

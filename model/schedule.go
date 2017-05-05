package model

import "github.com/songrgg/backeye/std"

// Project describes the API testing project
type Project struct {
	ID     int64  `gorm:"primary_key" json:"id"`
	Name   string `gorm:"type:varchar(256)" json:"name"`
	Desc   string `gorm:"type:varchar(1024)" json:"desc"`
	Status string `gorm:"type:varchar(64)" json:"status"`
	Tasks  []Task `json:"tasks"`
	std.TimeMixin
}

// Task sets the API testing task
type Task struct {
	ID        int64   `gorm:"primary_key" json:"id"`
	ProjectID int64   `gorm:"index" json:"project_id"`
	Name      string  `gorm:"type:varchar(256)" json:"name"`
	Type      string  `gorm:"type:varchar(64)" json:"type"`
	Status    string  `gorm:"type:varchar(64)" json:"status"`
	Desc      string  `gorm:"type:varchar(1024)" json:"desc"`
	CronSpec  string  `gorm:"type:varchar(256)" json:"cron_spec"`
	Watches   []Watch `gorm:"ForeignKey:TaskID;AssociationForeignKey:TaskID" json:"watches,omitempty"`
	std.TimeMixin
}

// Watch sets the tasks' API watch method
type Watch struct {
	ID         int64       `gorm:"primary_key" json:"id"`
	TaskID     int64       `gorm:"index" json:"task_id"`
	Name       string      `gorm:"type:varchar(256)" json:"name"`
	Desc       string      `gorm:"type:varchar(1024)" json:"desc"`
	Interval   int32       `json:"interval"`
	Timeout    int32       `json:"timeout"` // TODO: add timeout support
	Path       string      `gorm:"type:varchar(512)" json:"path"`
	Method     string      `gorm:"type:varchar(512)" json:"method"`
	Headers    string      `gorm:"type:longtext" json:"headers"`
	Assertions []Assertion `gorm:"ForeignKey:WatchID;AssociationForeignKey:ID" json:"assertions"`
	Variables  []Variable  `gorm:"ForeignKey:WatchID;AssociationForeignKey:ID" json:"variables"`
	std.TimeMixin
}

// Assertion indicates the watch's assertion logic
type Assertion struct {
	ID       int64  `gorm:"primary_key" json:"id"`
	WatchID  int64  `gorm:"index" json:"watch_id"`
	Type     string `gorm:"varchar(64)" json:"type"`
	Code     string `gorm:"varchar(32)" json:"code"`
	Source   string `gorm:"varchar(64)" json:"source"`
	Operator string `gorm:"varchar(64)" json:"operator"`
	Left     string `gorm:"varchar(128)" json:"left"`
	Right    string `gorm:"varchar(128)" json:"right"`
	Revision int32  `json:"revision"`
	std.TimeMixin
}

// Variable is
type Variable struct {
	ID      int64  `gorm:"primary_key"`
	WatchID int64  `gorm:"index"`
	Name    string `gorm:"varchar(128)"`
	Value   string `gorm:"type:longtext"`
}

// WatchResult indicates the watch's result
type WatchResult struct {
	ID               int64             `gorm:"primary_key" json:"id"`
	TaskID           int64             `gorm:"index" json:"task_id"`
	Status           string            `gorm:"type:varchar(32)" json:"status"`
	AssertionResults []AssertionResult `json:"assertion_results"`
	std.TimeMixin
}

// AssertionResult indicates the assertion result
type AssertionResult struct {
	ID                int64  `gorm:"primary_key" json:"id"`
	WatchResultID     int64  `gorm:"index" json:"watchresult_id"`
	AssertionID       int64  `gorm:"index" json:"assertion_id"`
	Status            string `gorm:"varchar(64)" json:"status"`
	Message           string `gorm:"varchar(2048)" json:"message"`
	ExecutionDuration int64  `json:"execution_duration"`
	std.TimeMixin
}

// TaskHealth indicates the task's health condition
type TaskHealth struct {
	Total   int `json:"total"`
	Success int `json:"success"`
}

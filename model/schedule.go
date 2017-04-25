package model

import "github.com/songrgg/backeye/std"

// Project describes the API testing project
type Project struct {
	ID     int64  `gorm:"primary_key"`
	Name   string `gorm:"type:varchar(256)"`
	Desc   string `gorm:"type:varchar(1024)"`
	Status string `gorm:"type:varchar(64)"`
	Tasks  []Task
	std.TimeMixin
}

// Task sets the API testing task
type Task struct {
	ID        int64   `gorm:"primary_key"`
	ProjectID int64   `gorm:"index"`
	Name      string  `gorm:"type:varchar(256)"`
	Type      string  `gorm:"type:varchar(64)"`
	Status    string  `gorm:"type:varchar(64)"`
	Desc      string  `gorm:"type:varchar(1024)"`
	CronSpec  string  `gorm:"type:varchar(256)"`
	Watches   []Watch `gorm:"ForeignKey:TaskID"`
	std.TimeMixin
}

// Watch sets the tasks' API watch method
type Watch struct {
	ID         int64  `gorm:"primary_key"`
	TaskID     int64  `gorm:"task_id"`
	Name       string `gorm:"type:varchar(256)"`
	Desc       string `gorm:"type:varchar(1024)"`
	Interval   int32
	Path       string `gorm:"type:varchar(512)"`
	Method     string `gorm:"type:varchar(512)"`
	Headers    string `gorm:"type:longtext"`
	Assertions []Assertion
	std.TimeMixin
}

// Assertion indicates the watch's assertion logic
type Assertion struct {
	ID       int64  `gorm:"primary_key"`
	WatchID  int64  `gorm:"index"`
	Type     string `gorm:"varchar(64)"`
	Code     string `gorm:"varchar(32)"`
	Timeout  int32  // TODO: add timeout support
	Source   string `gorm:"varchar(64)"`
	Operator string `gorm:"varchar(64)"`
	Left     string `gorm:"varchar(128)"`
	Right    string `gorm:"varchar(128)"`
	Revision int32
	std.TimeMixin
}

// WatchResult indicates the watch's result
type WatchResult struct {
	ID               int64  `gorm:"primary_key"`
	TaskID           int64  `gorm:"index"`
	Status           string `gorm:"type:varchar(32)"`
	AssertionResults []AssertionResult
	std.TimeMixin
}

// AssertionResult indicates the assertion result
type AssertionResult struct {
	ID                int64  `gorm:"primary_key"`
	WatchResultID     int64  `gorm:"index"`
	AssertionID       int64  `gorm:"index"`
	Status            string `gorm:"varchar(64)"`
	Message           string `gorm:"varchar(2048)"`
	ExecutionDuration int64
	std.TimeMixin
}

// TaskHealth indicates the task's health condition
type TaskHealth struct {
	Total   int `json:"total"`
	Success int `json:"success"`
}

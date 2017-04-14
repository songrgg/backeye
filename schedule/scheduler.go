package schedule

import (
	"context"
	"errors"
	"fmt"

	"github.com/robfig/cron"
	"github.com/songrgg/backeye/model"
	"github.com/songrgg/backeye/parser/json"
	"github.com/songrgg/backeye/std"
	"github.com/songrgg/backeye/target"
	"github.com/songrgg/backeye/watch"
)

type status int

const (
	// RUNNING indicates the cronjob is running
	RUNNING status = 1
	// STOPPED indicates the cronjob is stopped
	STOPPED
)

// Scheduler defines the api test's schedule rule
type Scheduler struct {
	schedules    map[string]*Schedule
	WatchResults chan watch.WatchResult
}

// Schedule includes a schedule target and cronjob
type Schedule struct {
	target *target.Task
	status status
	cron   *cron.Cron
}

var (
	// INSTANCE represents the single global schedule
	INSTANCE = newScheduler()
)

// LoadTasks loads target
func (sch *Scheduler) LoadTasks(targets []model.Task) error {
	parser := json.Parser{}
	for _, target := range targets {
		t2, err := parser.TranslateModel(&target)
		if err != nil {
			std.LogErrorc("scheduler", err, fmt.Sprintf("fail to parse target %s", target.Name))
			continue
		}
		if err := sch.Create(t2); err != nil {
			std.LogErrorc("scheduler", err, fmt.Sprintf("fail to create target %s", t2.Name))
			continue
		}
		if err := sch.Start(t2.Name); err != nil {
			std.LogErrorc("scheduler", err, fmt.Sprintf("fail to start target %s", t2.Name))
			continue
		}
	}
	return nil
}

func newScheduler() *Scheduler {
	return &Scheduler{
		schedules:    make(map[string]*Schedule),
		WatchResults: make(chan watch.WatchResult, 1000),
	}
}

// Create a schedule
func (sch *Scheduler) Create(t *target.Task) error {
	if _, err := sch.getSchedule(t.Name); err == nil {
		std.LogErrorc("mongo", nil, "schedule already exists")
		return nil
	}

	sch.schedules[t.Name] = &Schedule{
		target: t,
		cron:   parseCron(t, sch.WatchResults),
		status: STOPPED,
	}
	std.LogInfoc("mongo", fmt.Sprintf("schedule %s added", t.Name))
	return nil
}

// Start the specified schedule
func (sch *Scheduler) Start(name string) (err error) {
	var schedule *Schedule
	if schedule, err = sch.getSchedule(name); err != nil {
		std.LogErrorc("mongo", nil, fmt.Sprintf("schedule %s not exists", name))
		return err
	}
	schedule.status = RUNNING
	schedule.cron.Start()
	std.LogInfoc("mongo", fmt.Sprintf("schedule %s started", name))
	return nil
}

// Stop the specified schedule
func (sch *Scheduler) Stop(name string) (err error) {
	var schedule *Schedule
	if schedule, err = sch.getSchedule(name); err != nil {
		return err
	}
	schedule.status = STOPPED
	schedule.cron.Stop()
	return nil
}

// IsRunning indicates if the the schedule is running
func (sch *Scheduler) IsRunning(name string) (bool, error) {
	var (
		err      error
		schedule *Schedule
	)
	if schedule, err = sch.getSchedule(name); err != nil {
		return false, err
	}
	return schedule.status == RUNNING, nil
}

func (sch *Scheduler) getSchedule(name string) (*Schedule, error) {
	if schedule, ok := sch.schedules[name]; ok {
		return schedule, nil
	}
	return nil, errors.New("schedule not found")
}

func parseCron(t *target.Task, wr chan watch.WatchResult) *cron.Cron {
	c := cron.New()
	c.AddFunc(t.CronSpec, func() {
		results, err := t.Run(context.Background())
		if err != nil {
			std.LogErrorc("cron", err, "fail to run cron")
			return
		}

		for _, result := range results {
			select {
			case wr <- result:
			default:
				std.LogErrorc("watch_result", nil, "watch result channel is full")
				continue
			}
		}
	})
	return c
}

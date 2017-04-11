package schedule

import (
	"context"
	"errors"
	"fmt"

	"github.com/robfig/cron"
	"github.com/songrgg/backeye/hook"
	"github.com/songrgg/backeye/model"
	"github.com/songrgg/backeye/parser/json"
	"github.com/songrgg/backeye/std"
	"github.com/songrgg/backeye/target"
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
	schedules map[string]*Schedule
	hooks     []*hook.Hook
}

// Schedule includes a schedule target and cronjob
type Schedule struct {
	target *target.Target
	status status
	cron   *cron.Cron
}

var (
	// INSTANCE represents the single global schedule
	INSTANCE = newScheduler()
)

// LoadTargets loads target
func (sch *Scheduler) LoadTargets(targets []model.Target) error {
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

func (sch *Scheduler) AddHook(h *hook.Hook) {
	sch.hooks = append(sch.hooks, h)
}

func newScheduler() *Scheduler {
	return &Scheduler{
		schedules: make(map[string]*Schedule),
		hooks:     make([]*hook.Hook, 0),
	}
}

// Create a schedule
func (sch *Scheduler) Create(t *target.Target) error {
	if _, err := sch.getSchedule(t.Name); err == nil {
		std.LogErrorc("mongo", nil, "schedule already exists")
		return nil
	}

	sch.schedules[t.Name] = &Schedule{
		target: t,
		cron:   parseCron(t),
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

func parseCron(t *target.Target) *cron.Cron {
	c := cron.New()
	c.AddFunc(t.CronSpec, func() {
		results, err := t.Run(context.Background())
		if err != nil {
			std.LogErrorc("cron", err, "fail to run cron")
			return
		}

		for _, result := range results {
			fmt.Println("result: ", result.Assertions)
		}
	})
	return c
}

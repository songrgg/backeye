package scheduler

import (
	"context"

	"fmt"

	"encoding/json"

	"github.com/robfig/cron"
	"github.com/songrgg/backeye/dao"
	"github.com/songrgg/backeye/model"
	"github.com/songrgg/backeye/pkg/watcher"
)

type Scheduler struct {
	jobs map[int64]*cron.Cron
}

func (s *Scheduler) AddWatch(w *model.Watcher) error {
	wat, err := watcher.NewWatcher(w.Points)
	if err != nil {
		return err
	}

	c := cron.New()
	err = c.AddFunc(w.Cron, func() {
		as, _ := wat.Do(context.Background())
		fmt.Println(as)
		result, _ := json.Marshal(as)
		dao.NewAssertionResult(w.ID, true, result)
	})
	if err != nil {
		return err
	}

	s.jobs[w.ID] = c
	return nil
}

func (s *Scheduler) Stop(id int64) {
	s.jobs[id].Stop()
}

func (s *Scheduler) Start(id int64) {
	s.jobs[id].Start()
}

func NewScheduler() *Scheduler {
	scheduler := &Scheduler{}
	scheduler.jobs = make(map[int64]*cron.Cron)
	return scheduler
}

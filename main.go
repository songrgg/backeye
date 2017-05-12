// @APIVersion 1.0.0
// @APITitle Backeye
// @APIDescription Backeye usually works as expected.
// @Contact songrgg0.0@gmail.com
// @TermsOfServiceUrl http://google.com/
// @License MIT
// @BasePath /
package main

import (
	"time"

	"github.com/songrgg/backeye/dao"
	"github.com/songrgg/backeye/model"
	"github.com/songrgg/backeye/schedule"
	"github.com/songrgg/backeye/service"
	"github.com/songrgg/backeye/std"
)

func main() {
	std.InitLog(std.Config.Log)
	initSchedule()
	go watchResults()
	server.RunServer()
}

func initSchedule() {
	model.InitModel(std.Config.MySQL)

	all, err := dao.AllTasks()
	if err != nil {
		panic(err)
	}

	schedule.INSTANCE.LoadTasks(all)
}

func watchResults() {
	for true {
		select {
		case watchResult := <-schedule.INSTANCE.WatchResults:
			assertionResults := make([]model.AssertionResult, len(watchResult.Assertions))

			resultStatus := "success"
			for i, as := range watchResult.Assertions {
				now := time.Now()
				msg := ""
				if as.Error != nil {
					msg = as.Error.Error()
				}

				status := "failed"
				if as.Passed == true {
					status = "success"
				} else {
					resultStatus = "failed"
				}
				assertionResults[i] = model.AssertionResult{
					AssertionID:       as.AssertionID,
					Status:            status,
					ExecutionDuration: as.ExecutionDuration.Nanoseconds(),
					Message:           msg,
					TimeMixin: std.TimeMixin{
						CreatedAt: now,
						UpdatedAt: now,
					},
				}
			}

			err := dao.NewWatchResult(&model.WatchResult{
				TaskID:           watchResult.TaskID,
				Status:           resultStatus,
				AssertionResults: assertionResults,
			})
			if err != nil {
				std.LogErrorc("mysql", err, "fail to store watch result")
			}
		}
	}
}

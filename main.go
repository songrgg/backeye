package main

import (
	"github.com/songrgg/backeye/common"
	"github.com/songrgg/backeye/dao"
	"github.com/songrgg/backeye/model"
	"github.com/songrgg/backeye/schedule"
	"github.com/songrgg/backeye/server"
	"github.com/songrgg/backeye/std"
)

func main() {
	std.InitLog(common.Config.Log)
	initSchedule()
	go watchResults()
	server.RunServer()
}

func initSchedule() {
	model.InitModel(common.Config.MySQL)

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
				msg := ""
				if as.Error != nil {
					msg = as.Error.Error()
				}

				status := "failed"
				if as.Success == true {
					status = "success"
				} else {
					resultStatus = "failed"
				}
				assertionResults[i] = model.AssertionResult{
					Status:  status,
					Message: msg,
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

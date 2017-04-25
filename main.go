package main

import (
	"fmt"

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
	// go watchResults()
	server.RunServer()
}

func initSchedule() {
	model.InitModel(common.Config.MySQL)
	// dao.InitMongo()

	all, err := dao.AllTasks()
	if err != nil {
		panic(err)
	}
	fmt.Println(all)

	schedule.INSTANCE.LoadTasks(all)
}

// func watchResults() {
// 	for true {
// 		select {
// 		case watchResult := <-schedule.INSTANCE.WatchResults:
// 			fmt.Println("results: ", watchResult.Assertions)
// 			assertionResults := make([]model.AssertionResult, len(watchResult.Assertions))
// 			for i, as := range watchResult.Assertions {
// 				msg := ""
// 				if as.Error != nil {
// 					msg = as.Error.Error()
// 				}
// 				assertionResults[i] = model.AssertionResult{
// 					Success: as.Success,
// 					Message: msg,
// 				}
// 			}
// 			err := dao.NewWatchResult(&model.WatchResult{
// 				TaskName:      watchResult.TaskName,
// 				WatchName:     watchResult.WatchName,
// 				ExecutionTime: watchResult.ExecutionTime.Unix(),
// 				Assertions:    assertionResults,
// 			})
// 			if err != nil {
// 				std.LogErrorc("mongo", err, "fail to store watch result")
// 			}
// 		}
// 	}
// }

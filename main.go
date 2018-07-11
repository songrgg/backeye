// @APIVersion 1.0.0
// @APITitle Backeye
// @APIDescription Backeye usually works as expected.
// @Contact songrgg0.0@gmail.com
// @TermsOfServiceUrl http://google.com/
// @License MIT
// @BasePath /
package main

import (
	"fmt"

	"github.com/songrgg/backeye/dao"
	"github.com/songrgg/backeye/model"
	"github.com/songrgg/backeye/scheduler"
	"github.com/songrgg/backeye/service"
	"github.com/songrgg/backeye/std"
)

func main() {
	std.InitLog(std.Config.Log)
	initSchedule()
	server.RunServer()
}

func initSchedule() {
	model.InitModel(std.Config.MySQL)

	all, err := dao.GetWatchers()
	if err != nil {
		panic(err)
	}
	fmt.Println(all)

	sched := scheduler.NewScheduler()

	for _, w := range all {
		err := sched.AddWatch(&w)
		if err != nil {
			panic(err)
		}

		sched.Start(w.ID)
	}
}

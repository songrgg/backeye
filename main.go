package main

import (
	"github.com/songrgg/backeye/common"
	"github.com/songrgg/backeye/dao"
	"github.com/songrgg/backeye/schedule"
	"github.com/songrgg/backeye/server"
	"github.com/songrgg/backeye/std"
)

func main() {
	std.InitLog(common.Config.Log)
	initSchedule()
	server.RunServer()
}

func initSchedule() {
	dao.InitMongo()

	all, err := dao.All()
	if err != nil {
		panic(err)
	}
	schedule.INSTANCE.LoadTargets(all)
}

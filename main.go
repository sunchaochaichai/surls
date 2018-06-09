package main

import (
	_ "surls/global"
	"surls/servers"
	"github.com/sirupsen/logrus"
	"surls/lib"
	"surls/global"
)

func main() {

	lib.SetPid(global.ProjectRealPath + "/runtime/pid")

	go func() {
		servers.RunHttpServer()
	}()

	go func() {
		servers.RunMetricsServer()
	}()

	go func() {
		servers.RunDebug()
	}()

	logrus.Info("done", servers.RunGrpcServer())
}

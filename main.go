package main

import (
	_ "surls/global"
	"surls/servers"
	"surls/lib"
	"github.com/sirupsen/logrus"
)

func main() {

	lib.SetPid("pid")

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

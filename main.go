package main

import (
	"surls/servers"
	"surls/global"
	"surls/lib"
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

	global.Logger.Log("done", servers.RunGrpcServer())
}

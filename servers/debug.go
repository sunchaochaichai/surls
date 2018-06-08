package servers

import (
	"net/http"
	_ "net/http/pprof"
	"surls/cli"
	"strings"
	_ "github.com/mkevac/debugcharts"
	"surls/global"
)

func RunDebug() {
	global.Logger.Log(
		"transport",
		"debug",
		"addr", cli.Params.DebugAddr,
		"path",
		strings.Join([]string{
			"/debug/pprof/",
			"/debug/pprof/cmdline",
			"/debug/pprof/profile",
			"/debug/pprof/symbol",
			"/debug/pprof/trace",
			"/debug/charts/",
		}, ","),
	)

	global.Logger.Log(
		"err",
		http.ListenAndServe(cli.Params.DebugAddr, nil),
	)
}

package servers

import (
	"net/http"
	_ "net/http/pprof"
	"surls/cli"
	_ "github.com/mkevac/debugcharts"
	"github.com/sirupsen/logrus"
	"strings"
	"log"
	"fmt"
)

func RunDebug() {
	info := make(logrus.Fields)
	info["transport"] = "debug"
	info["addr"] = cli.Params.DebugAddr
	info["path"] = strings.Join([]string{
		"/debug/pprof/",
		"/debug/pprof/cmdline",
		"/debug/pprof/profile",
		"/debug/pprof/symbol",
		"/debug/pprof/trace",
		"/debug/charts/",
	}, ",")
	log.Println(
		fmt.Sprintf(
			"%s , %s=%s , %s=%s , %s=%s",
			"debug open...",
			"transport", "debug",
			"addr", cli.Params.DebugAddr,
			"path", strings.Join([]string{
				"/debug/pprof/",
				"/debug/pprof/cmdline",
				"/debug/pprof/profile",
				"/debug/pprof/symbol",
				"/debug/pprof/trace",
				"/debug/charts/",
			}, ","),
		),
	)

	log.Fatalln(
		"err",
		http.ListenAndServe(cli.Params.DebugAddr, nil),
	)
}

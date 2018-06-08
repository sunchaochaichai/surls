package servers

import (
	"net/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"surls/cli"
	"fmt"
	"surls/global"
)

func RunMetricsServer() {
	mux := http.NewServeMux()
	mux.HandleFunc(
		cli.Params.MetricsPath,
		promhttp.Handler().ServeHTTP,
	)

	global.Logger.Log(
		"transport",
		"instrumenting",
		"adapter",
		"prometheus",
		"addr",
		fmt.Sprintf(
			"%s%s",
			cli.Params.MetricsAddr,
			cli.Params.MetricsPath,
		),
	)

	global.Logger.Log(
		"err",
		http.ListenAndServe(cli.Params.MetricsAddr, mux),
	)
}

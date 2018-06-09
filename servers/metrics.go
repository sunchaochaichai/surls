package servers

import (
	"net/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"surls/cli"
	"fmt"
	"log"
)

func RunMetricsServer() {
	mux := http.NewServeMux()
	mux.HandleFunc(
		cli.Params.MetricsPath,
		promhttp.Handler().ServeHTTP,
	)

	log.Println(
		fmt.Sprintf(
			"%s , %s=%s , %s=%s , %s=%s%s ,",
			"metrics running...",
			"transport", "instrumenting",
			"adapter", "prometheus",
			"addr", cli.Params.MetricsAddr, cli.Params.MetricsPath,
		),
	)

	log.Fatalln(
		"metrics error:",
		http.ListenAndServe(cli.Params.MetricsAddr, mux),
	)
}

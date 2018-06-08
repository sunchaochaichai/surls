package surlssvc

import (
	"surls/svc/surlssvc/interfaces"
	"surls/svc/surlssvc/middlewares/mw_svc"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"context"
)

//实现服务
type SUrls struct {
}

func (SUrls) Get(ctx context.Context, s string) (output string, err error) {
	output = s
	return
}

func (SUrls) Set(ctx context.Context, s string) (output string, err error) {
	output = s
	return
}

func SUrlsSvc() interfaces.SUrlsInf {
	var svc interfaces.SUrlsInf
	svc = SUrls{}

	//日志
	//svc = mw_svc.LoggingMiddleware{global.Logger, svc}

	//prometheus 中间件
	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{})
	svc = mw_svc.Instrumenting(requestCount, requestLatency, countResult)(svc)

	return svc
}

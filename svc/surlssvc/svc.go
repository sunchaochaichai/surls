package surlssvc

import (
	"surls/svc/surlssvc/interfaces"
	"surls/svc/surlssvc/middlewares/mw_svc"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"context"
	"surls/global"
	"github.com/go-siris/siris/core/errors"
	"surls/utils"
)

//实现服务
type SUrls struct {
}

func (SUrls) Get(ctx context.Context, shortUrl string) (entity interfaces.SurlEntity, err error) {
	sourceUrl := global.Redis.Get(shortUrl).Val()
	entity.ShortUrl = shortUrl
	if sourceUrl == "" {
		err = errors.New("domain not found")
		return
	}

	entity.SourceUrl = sourceUrl
	return
}

func (SUrls) Set(ctx context.Context, sourceUrl string) (entity interfaces.SurlEntity, err error) {
	entity.SourceUrl = sourceUrl
	entity.ShortUrl = utils.StrMd5(sourceUrl)

	err = global.Redis.Set(
		entity.ShortUrl,
		entity.SourceUrl,
		0,
	).Err()

	return
}

func SUrlsSvc() interfaces.SUrlsInf {
	var svc interfaces.SUrlsInf
	svc = SUrls{}

	//日志
	svc = mw_svc.LoggingMiddleware{svc}

	//prometheus 中间件
	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "svc_status",
		Subsystem: "susls",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "svc_status",
		Subsystem: "susls",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "svc_status",
		Subsystem: "surls",
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{})
	svc = mw_svc.Instrumenting(requestCount, requestLatency, countResult)(svc)

	return svc
}

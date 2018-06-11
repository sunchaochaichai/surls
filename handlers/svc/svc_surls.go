package svc

import (
	"context"
	"surls/global"
	"errors"
	"surls/utils"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/go-kit/kit/metrics/prometheus"
)

func New(logger global.Logger) SUrlsInf {
	var svc SUrlsInf
	svc = SUrls{}

	//日志
	svc = LoggingMiddleware(logger)(svc)

	//prometheus 中间件
	fieldKeys := []string{"method", "error"}
	requestCount := prometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "svc_status",
		Subsystem: "susls",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := prometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "svc_status",
		Subsystem: "susls",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	countResult := prometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "svc_status",
		Subsystem: "surls",
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{})
	svc = InstrumentingMiddleware(requestCount, requestLatency, countResult)(svc)

	return svc
}


//定义服务接口
type SUrlsInf interface {
	Get(ctx context.Context, shortUrl string) (SurlEntity, error)
	Set(ctx context.Context, sourceUrl string) (SurlEntity, error)
}

type SurlEntity struct {
	SourceUrl string
	ShortUrl  string
}

//定义中间件接口
type SUrlsMiddleware func(SUrlsInf) SUrlsInf

//实现服务
type SUrls struct {
}

func (SUrls) Get(ctx context.Context, shortUrl string) (entity SurlEntity, err error) {
	sourceUrl := global.Redis.Get(shortUrl).Val()
	entity.ShortUrl = shortUrl
	if sourceUrl == "" {
		err = errors.New("domain not found")
		return
	}

	entity.SourceUrl = sourceUrl
	return
}

func (SUrls) Set(ctx context.Context, sourceUrl string) (entity SurlEntity, err error) {
	entity.SourceUrl = sourceUrl
	entity.ShortUrl = utils.StrMd5(sourceUrl)

	err = global.Redis.Set(
		entity.ShortUrl,
		entity.SourceUrl,
		0,
	).Err()

	return
}
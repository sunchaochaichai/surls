package endpoints

import (
	"github.com/go-kit/kit/endpoint"
	"context"
	"surls/handlers/surls/svc"
	"surls/pb"
	//"github.com/go-kit/kit/metrics"
	"surls/global"
	"github.com/go-kit/kit/metrics"
	"time"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/sony/gobreaker"
	"github.com/go-kit/kit/ratelimit"
	"golang.org/x/time/rate"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/go-kit/kit/metrics/prometheus"
)

func New(
	svc svc.SUrlsInf,
	logger global.Logger,
//otTracer stdopentracing.Tracer,
//zipkinTracer *stdzipkin.Tracer,
) Surls {
	var duration metrics.Histogram
	{
		// Endpoint-level metrics.
		duration = prometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "business",
			Subsystem: "surls",
			Name:      "request_duration_seconds",
			Help:      "Request duration in seconds.",
		}, []string{"method", "success"})
	}

	//	//熔断中间件
	//	hystrix.ConfigureCommand("set", hystrix.CommandConfig{
	//		Timeout:               60000,
	//		ErrorPercentThreshold: 10,
	//		MaxConcurrentRequests: 20,
	//	})
	//
	//	setEndpoint = Hystrix("surls_set")(setEndpoint)

	var getEndpoint endpoint.Endpoint
	{
		getEndpoint = MakeGetEndpoint(svc)
		getEndpoint = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 100))(getEndpoint)
		getEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:        "surls_get",
			MaxRequests: 10000,
			Interval:    5 * time.Second,
			Timeout:     60 * time.Second,
		}))(getEndpoint)
		//getEndpoint = opentracing.TraceServer(otTracer, "Sum")(sumEndpoint)
		//getEndpoint = zipkin.TraceEndpoint(zipkinTracer, "Sum")(sumEndpoint)
		getEndpoint = LoggingMiddleware(global.Log)(getEndpoint)
		getEndpoint = InstrumentingMiddleware(duration.With("method", "surls_get"))(getEndpoint)
	}
	var setEndpoint endpoint.Endpoint
	{
		setEndpoint = MakeSetEndpoint(svc)
		setEndpoint = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 100))(setEndpoint)
		setEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:        "surls_get",
			MaxRequests: 5000,
			Interval:    5 * time.Second,
			Timeout:     60 * time.Second,
		}))(setEndpoint)
		//setEndpoint = opentracing.TraceServer(otTracer, "Concat")(setEndpoint)
		//setEndpoint = zipkin.TraceEndpoint(zipkinTracer, "Concat")(setEndpoint)
		setEndpoint = LoggingMiddleware(global.Log)(setEndpoint)
		setEndpoint = InstrumentingMiddleware(duration.With("method", "surls_set"))(setEndpoint)
	}
	return Surls{
		GetEndpoint: getEndpoint,
		SetEndpoint: setEndpoint,
	}
}

type Surls struct {
	GetEndpoint endpoint.Endpoint
	SetEndpoint endpoint.Endpoint
}

func (this Surls) Get(ctx context.Context, shortUrl string) (entity svc.SurlEntity, err error) {
	entity.ShortUrl = shortUrl
	resp, err := this.GetEndpoint(ctx, &pb.GetReq{Url: shortUrl})
	if err != nil {
		return entity, err
	}

	response := resp.(pb.GetResp)

	entity.ShortUrl = response.Data

	return
}

func (this Surls) Set(ctx context.Context, sourceUrl string) (entity svc.SurlEntity, err error) {
	entity.SourceUrl = sourceUrl
	req := &pb.SetReq{Url: sourceUrl}
	resp, err := this.SetEndpoint(ctx, req)
	if err != nil {
		return
	}
	response := resp.(pb.SetResp)

	entity.ShortUrl = response.Data.ShortUrl
	return
}

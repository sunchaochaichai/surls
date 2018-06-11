package svc

import (
	"github.com/go-kit/kit/metrics"
	"context"
	"time"
	"fmt"
)

func InstrumentingMiddleware(
	requestCount metrics.Counter,
	requestLatency metrics.Histogram,
	countResult metrics.Histogram,
) SUrlsMiddleware {
	return func(next SUrlsInf) SUrlsInf {
		return Instrmw{
			requestCount,
			requestLatency,
			countResult,
			next,
		}
	}
}

type Instrmw struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	countResult    metrics.Histogram
	SUrlsInf
}

func (this Instrmw) Set(ctx context.Context,s string) (entity SurlEntity, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "set", "error", fmt.Sprint(err != nil)}
		this.requestCount.With(lvs...).Add(1)
		this.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	entity, err = this.SUrlsInf.Set(ctx,s)
	return
}

func (this Instrmw) Get(ctx context.Context,s string) (entity SurlEntity, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "get", "error", "false"}
		this.requestCount.With(lvs...).Add(1)
		this.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	entity,err = this.SUrlsInf.Get(ctx,s)
	return
}

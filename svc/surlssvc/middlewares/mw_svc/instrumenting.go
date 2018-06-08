package mw_svc

import (
	"time"
	"github.com/go-kit/kit/metrics"
	"fmt"
	"surls/svc/surlssvc/interfaces"
	"context"
)

func Instrumenting(
	requestCount metrics.Counter,
	requestLatency metrics.Histogram,
	countResult metrics.Histogram,
) interfaces.SUrlsMiddleware {
	return func(next interfaces.SUrlsInf) interfaces.SUrlsInf {
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
	interfaces.SUrlsInf
}

func (this Instrmw) Set(ctx context.Context,s string) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "upper", "error", fmt.Sprint(err != nil)}
		this.requestCount.With(lvs...).Add(1)
		this.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = this.SUrlsInf.Set(ctx,s)
	return
}

func (this Instrmw) Get(ctx context.Context,s string) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "count", "error", "false"}
		this.requestCount.With(lvs...).Add(1)
		this.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output,err = this.SUrlsInf.Get(ctx,s)
	return
}
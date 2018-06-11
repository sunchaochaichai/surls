package endpoints

import (
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/endpoint"
	"time"
	"context"
	"fmt"
)

func InstrumentingMiddleware(duration metrics.Histogram) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				duration.With(
					"success",
					fmt.Sprint(err == nil)).
					Observe(time.Since(begin).
					Seconds(),
				)
			}(time.Now())
			return next(ctx, request)
		}
	}
}

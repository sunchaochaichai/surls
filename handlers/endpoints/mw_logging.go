package endpoints

import (
	"github.com/go-kit/kit/endpoint"
	"time"
	"surls/global"
	"context"
)

func LoggingMiddleware(logger global.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {

			defer func(begin time.Time) {
				logger.Log("transport_error:", err, "took:", time.Since(begin))
			}(time.Now())

			return next(ctx, request)
		}
	}
}
package mw_endpoint

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/go-kit/kit/endpoint"
	"surls/pb"
	"surls/lib/resp_errors"
)

func SetHystrix(commandName string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			var resp interface{}
			err = hystrix.Do(commandName, func() (err error) {
				resp, err = next(ctx, request)
				return err
			}, nil)

			if err != nil {
				return pb.SetResp{
					Code: resp_errors.ERROR_TOO_MANY_CONNECTIONS.Code,
					Msg:  resp_errors.ERROR_TOO_MANY_CONNECTIONS.Msg,
				}, nil
			}

			return resp, nil
		}
	}
}

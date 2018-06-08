package endpoints

import (
	"github.com/go-kit/kit/endpoint"
	"surls/svc/surlssvc/middlewares/mw_endpoint"
	"github.com/afex/hystrix-go/hystrix"
	"surls/svc/surlssvc/interfaces"
	"surls/lib/resp_errors"
	"context"
	"surls/pb"
)

func Set(svc interfaces.SUrlsInf) endpoint.Endpoint {
	var setEndpoint endpoint.Endpoint
	setEndpoint = MakeSetEndpoint(svc)

	////熔断中间件
	hystrix.ConfigureCommand("upper", hystrix.CommandConfig{
		Timeout:               60000,
		ErrorPercentThreshold: 10,
		MaxConcurrentRequests: 20,
	})

	setEndpoint = mw_endpoint.SetHystrix("set")(setEndpoint)

	return setEndpoint
}

//定义端点
func MakeSetEndpoint(svc interfaces.SUrlsInf) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		req := request.(*pb.SetReq)
		sourceUrl, err := svc.Set(ctx, req.Url)
		if err != nil {
			resp = pb.GetResp{
				Code: resp_errors.ERROR_PARAMS_ERROR.Code,
				Msg:  resp_errors.ERROR_PARAMS_ERROR.Msg,
			}
			return
		}
		resp = pb.SetResp{
			Code: resp_errors.SUCCESS.Code,
			Msg:  resp_errors.SUCCESS.Msg,
			Data: &pb.SetRespData{
				SourceUrl: sourceUrl,
				ShortUrl:  sourceUrl,
			},
		}
		return
	}
}

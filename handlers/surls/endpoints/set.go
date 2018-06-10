package endpoints

import (
	"github.com/go-kit/kit/endpoint"
	"surls/handlers/surls/svc"
	"surls/lib/resp_errors"
	"context"
	"surls/pb"
)

//定义端点
func MakeSetEndpoint(svc svc.SUrlsInf) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		req := request.(*pb.SetReq)
		entity, err := svc.Set(ctx, req.Url)
		if err != nil {
			resp = pb.GetResp{
				Code: resp_errors.ERROR_PARAMS_ERROR.Code,
				Msg:  err.Error(),
			}
			return
		}
		resp = pb.SetResp{
			Code: resp_errors.SUCCESS.Code,
			Msg:  resp_errors.SUCCESS.Msg,
			Data: &pb.SetRespData{
				SourceUrl: entity.SourceUrl,
				ShortUrl:  entity.ShortUrl,
			},
		}
		return
	}
}

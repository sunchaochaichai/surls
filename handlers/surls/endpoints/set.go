package endpoints

import (
	"github.com/go-kit/kit/endpoint"
	"surls/handlers/surls/svc"
	"context"
	"surls/pb"
	"surls/global"
)

//定义端点
func MakeSetEndpoint(svc svc.SUrlsInf) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		req := request.(*pb.SetReq)
		entity, err := svc.Set(ctx, req.Url)
		if err != nil {
			resp = pb.GetResp{
				Code: global.ERROR_PARAMS_ERROR.Code,
				Msg:  err.Error(),
			}
			return
		}
		resp = pb.SetResp{
			Code: global.SUCCESS.Code,
			Msg:  global.SUCCESS.Msg,
			Data: &pb.SetRespData{
				SourceUrl: entity.SourceUrl,
				ShortUrl:  entity.ShortUrl,
			},
		}
		return
	}
}

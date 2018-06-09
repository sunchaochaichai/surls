package endpoints

import (
	"github.com/go-kit/kit/endpoint"
	"surls/svc/surlssvc/interfaces"
	"surls/lib/resp_errors"
	"context"
	"surls/pb"
)

func Get(svc interfaces.SUrlsInf) endpoint.Endpoint {
	return MakeGetEndpoint(svc)
}

func MakeGetEndpoint(svc interfaces.SUrlsInf) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
		req := request.(*pb.GetReq)
		entity, svcErr := svc.Get(ctx, req.Url)
		if svcErr != nil {
			resp = pb.GetResp{
				Code: resp_errors.ERROR_PARAMS_ERROR.Code,
				Msg:  svcErr.Error(),
			}
			return
		}
		resp = pb.GetResp{
			Code: resp_errors.SUCCESS.Code,
			Msg:  resp_errors.SUCCESS.Msg,
			Data: entity.SourceUrl,
		}
		return
	}
}

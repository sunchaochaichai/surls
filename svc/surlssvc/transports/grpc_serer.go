package transports

import (
	"context"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"surls/svc/surlssvc/endpoints"
	"surls/pb"
	"surls/svc/surlssvc"
)

type GrpcService struct {
	GetHandler grpctransport.Handler
	SetHandler grpctransport.Handler
}

func (this *GrpcService) Get(ctx context.Context, req *pb.GetReq) (*pb.GetResp, error) {
	_, resp, err := this.GetHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.GetResp), nil
}

func (this *GrpcService) Set(ctx context.Context, req *pb.SetReq) (*pb.SetResp, error) {
	_, resp, err := this.SetHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.SetResp), nil
}

func NewSUrlsGrpcServer() *GrpcService {

	svc := surlssvc.SUrlsSvc()

	getHandler := grpctransport.NewServer(
		endpoints.Get(svc),
		GrpcDecodeGetReq,
		GrpcEncodeGetResp,
	)

	setHandler := grpctransport.NewServer(
		endpoints.Set(svc),
		GrpcDecodeSetReq,
		GrpcEncodeSetResp,
	)

	return &GrpcService{
		GetHandler: getHandler,
		SetHandler: setHandler,
	}
}

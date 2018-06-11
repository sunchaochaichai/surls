package transports

import (
	"context"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"surls/handlers/endpoints"
	"surls/pb"
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

func NewSUrlsGrpcServer(endpoints endpoints.Surls) *GrpcService {

	//zipkinV2URL := "http://localhost:9411/api/v2/spans"
	//
	//var tracer = stdopentracing.GlobalTracer()
	//
	//var zipkinTracer *zipkingo.Tracer
	//{
	//	var (
	//		err           error
	//		hostPort      = "localhost:7070"
	//		serviceName   = "addsvc"
	//		useNoopTracer = false
	//		reporter      = zipkinhttp.NewReporter(zipkinV2URL)
	//	)
	//	defer reporter.Close()
	//	zEP, _ := zipkingo.NewEndpoint(serviceName, hostPort)
	//	zipkinTracer, err = zipkingo.NewTracer(
	//		reporter,
	//		zipkingo.WithLocalEndpoint(zEP),
	//		zipkingo.WithNoopTracer(useNoopTracer),
	//	)
	//	if err != nil {
	//		log.Fatalln("err", err)
	//	}
	//	log.Println("tracer", "Zipkin", "type", "Native", "URL", zipkinV2URL)
	//}
	//
	//zipkinServer := zipkin.GRPCServerTrace(zipkinTracer)
	//
	//logger := gokitlog.NewLogfmtLogger(os.Stderr)
	//
	//options := []grpctransport.ServerOption{
	//	zipkinServer,
	//	grpctransport.ServerBefore(
	//		opentracing.GRPCToContext(
	//			tracer,
	//			"Sum",
	//			logger,
	//		)),
	//}

	getHandler := grpctransport.NewServer(
		endpoints.GetEndpoint,
		GrpcDecodeGetReq,
		GrpcEncodeGetResp,
		//options...,
	)

	setHandler := grpctransport.NewServer(
		endpoints.SetEndpoint,
		GrpcDecodeSetReq,
		GrpcEncodeSetResp,
	)

	return &GrpcService{
		GetHandler: getHandler,
		SetHandler: setHandler,
	}
}

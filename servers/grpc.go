package servers

import (
	"os"
	"google.golang.org/grpc"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"net"
	"surls/cli"
	"os/signal"
	"syscall"
	"surls/pb"
	"surls/handlers/surls/transports"
	"surls/handlers/surls/endpoints"
	"surls/handlers/surls/svc"
	"surls/global"
	"log"
)

func RunGrpcServer() error {

	grpcListener, err := net.Listen("tcp", cli.Params.GrpcServAddr)
	if err != nil {
		log.Fatalln("transport=gRPC , err=", err)
	}

	svc := svc.New(global.Log)

	endpoints := endpoints.New(svc, global.Log)

	grpcServer := transports.NewSUrlsGrpcServer(endpoints)

	baseServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))

	pb.RegisterSUrlsServer(
		baseServer,
		grpcServer,
	)

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		log.Println("grpc server exit...", <-c)
		baseServer.GracefulStop()
	}()

	log.Println("grpc server running... , transport=gRPC , addr=", cli.Params.GrpcServAddr)

	return baseServer.Serve(grpcListener)
}

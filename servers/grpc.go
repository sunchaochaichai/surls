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
	"surls/svc/surlssvc/transports"
	"log"
)

func RunGrpcServer() error {

	grpcListener, err := net.Listen("tcp", cli.Params.GrpcServAddr)
	if err != nil {
		log.Fatalln("transport=gRPC , err=", err)
	}

	server := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))

	handler := transports.NewSUrlsGrpcServer()

	pb.RegisterSUrlsServer(
		server,
		handler,
	)

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		log.Println("grpc server exit...", <-c)
		server.GracefulStop()
	}()

	log.Println("grpc server running... , transport=gRPC , addr=",cli.Params.GrpcServAddr)

	return server.Serve(grpcListener)
}

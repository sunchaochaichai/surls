package servers

import (
	"os"
	"google.golang.org/grpc"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"net"
	"surls/cli"
	"os/signal"
	"syscall"
	"surls/global"
	"surls/pb"
	"surls/svc/surlssvc/transports"
)

func RunGrpcServer() error {

	grpcListener, err := net.Listen("tcp", cli.Params.GrpcServAddr)
	if err != nil {
		global.Logger.Log("transport", "gRPC", "during", "Listen", "err", err)
		os.Exit(1)
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
		global.Logger.Log("exit...", <-c)
		server.GracefulStop()
	}()

	global.Logger.Log("transport", "gRPC", "addr", cli.Params.GrpcServAddr)
	return server.Serve(grpcListener)

}

package servers

import (
	"net/http"
	"context"
	"google.golang.org/grpc"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"surls/pb"
	"surls/cli"
	"os"
	"surls/global"
)

func RunHttpServer() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterSUrlsHandlerFromEndpoint(
		ctx,
		mux,
		"localhost"+cli.Params.GrpcServAddr,
		opts,
	)

	if err != nil {
		global.Logger.Log("err", err)
		os.Exit(2)
	}

	global.Logger.Log(
		"transport",
		"http",
		"addr",
		cli.Params.HttpServAddr,
		"proxy-grpc",
		cli.Params.GrpcServAddr,
	)

	global.Logger.Log(
		"err",
		http.ListenAndServe(cli.Params.HttpServAddr, mux),
	)
}

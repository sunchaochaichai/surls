package servers

import (
	"net/http"
	"context"
	"google.golang.org/grpc"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"surls/pb"
	"surls/cli"
	"log"
	"fmt"
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
		log.Fatalln("http server error:", err)
	}

	log.Println(
		fmt.Sprintf(
			"%s , %s=%s , %s=%s, %s=%s",
			"http proxy running...",
			"transport", "http",
			"addr", cli.Params.HttpServAddr,
			"grpc proxy", cli.Params.GrpcServAddr,
		),
	)

	log.Fatalln(
		"Http Server Error:",
		http.ListenAndServe(cli.Params.HttpServAddr, mux),
	)
}

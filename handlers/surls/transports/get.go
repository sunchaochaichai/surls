package transports

import (
	"golang.org/x/net/context"
	"surls/pb"
)

func GrpcDecodeGetReq(ctx context.Context, req interface{}) (interface{}, error) {
	getReq := req.(*pb.GetReq)
	return getReq, nil
}

func GrpcEncodeGetResp(ctx context.Context, resp interface{}) (interface{}, error) {
	getResp := resp.(pb.GetResp)
	return &getResp, nil
}

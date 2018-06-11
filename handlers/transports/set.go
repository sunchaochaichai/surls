package transports

import (
	"golang.org/x/net/context"
	"surls/pb"
)

func GrpcDecodeSetReq(_ context.Context, req interface{}) (interface{}, error) {
	setReq := req.(*pb.SetReq)
	return setReq, nil
}

func GrpcEncodeSetResp(_ context.Context, resp interface{}) (interface{}, error) {
	setResp := resp.(pb.SetResp)
	return &setResp, nil
}
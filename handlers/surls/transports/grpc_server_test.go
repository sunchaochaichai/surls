package transports

import (
	"testing"
	"surls/pb"
	"context"
	"reflect"
	"surls/handlers/surls/svc"
	"surls/handlers/surls/endpoints"
	"surls/global"
)

var surlsGrpcService = NewSUrlsGrpcServer(endpoints.New(
	svc.New(global.Log),
	global.Log,
))

func TestSurlsSet(t *testing.T) {
	tests := []struct {
		in   string
		want pb.SetResp
	}{
		{
			in: "http://a.com",
			want: pb.SetResp{
				Code: 200,
				Msg:  "success",
				Data: &pb.SetRespData{
					SourceUrl: "http://a.com",
					ShortUrl:  "388495550d157e5c5d299b7ecfeb1c2d",
				},
			},
		},
		{
			in: "http://b.com",
			want: pb.SetResp{
				Code: 200,
				Msg:  "success",
				Data: &pb.SetRespData{
					SourceUrl: "http://b.com",
					ShortUrl:  "715af3fd14408eda374ef3bcb23c10b6",
				},
			},
		},
	}

	for _, tt := range tests {
		req := &pb.SetReq{Url: tt.in}
		resp, err := surlsGrpcService.Set(context.Background(), req)

		if err != nil {
			t.Error("surls set test failed", err)
		}

		r := pb.SetResp{
			Code: resp.Code,
			Msg:  resp.Msg,
			Data: resp.Data,
		}

		if !reflect.DeepEqual(r.Data, tt.want.Data) {
			t.Errorf(
				"surls get(%s) = %s , wanted %s",
				tt.in,
				resp.String(),
				tt.want.String(),
			)
		}
	}
}

func BenchmarkSurlsSet(b *testing.B) {
	b.ResetTimer()
	req := &pb.SetReq{Url: "http://www.baidu.com"}
	for i := 0; i < b.N; i++ {
		surlsGrpcService.Set(context.Background(), req)
	}
}

func TestSurlsGet(t *testing.T) {

	tests := []struct {
		in   string
		want pb.GetResp
	}{
		{
			in: "388495550d157e5c5d299b7ecfeb1c2d",
			want: pb.GetResp{
				Code: 200,
				Msg:  "success",
				Data: "http://a.com",
			},
		},
		{
			in: "715af3fd14408eda374ef3bcb23c10b6",
			want: pb.GetResp{
				Code: 200,
				Msg:  "success",
				Data: "http://b.com",
			},
		},
	}

	for _, tt := range tests {
		req := &pb.GetReq{Url: tt.in}
		resp, err := surlsGrpcService.Get(context.Background(), req)

		if err != nil {
			t.Error("surls get test failed", err)
		}

		r := pb.GetResp{
			Code: resp.Code,
			Msg:  resp.Msg,
			Data: resp.Data,
		}

		if !reflect.DeepEqual(r.Data, tt.want.Data) {
			t.Errorf(
				"surls get(%s) = %s , wanted %s",
				tt.in,
				resp.String(),
				tt.want.String(),
			)
		}
	}
}

func BenchmarkSurlsGet(b *testing.B) {
	b.ResetTimer()
	req := &pb.GetReq{Url: "http://www.baidu.com"}
	for i := 0; i < b.N; i++ {
		surlsGrpcService.Get(context.Background(), req)
	}
}

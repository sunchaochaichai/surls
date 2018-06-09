package interfaces

import (
	"context"
)

type SurlEntity struct {
	SourceUrl string
	ShortUrl  string
}

//定义服务接口
type SUrlsInf interface {
	Get(context.Context, string) (SurlEntity, error)
	Set(context.Context, string) (SurlEntity, error)
}

//定义中间件接口
type SUrlsMiddleware func(SUrlsInf) SUrlsInf

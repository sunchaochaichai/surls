package interfaces

import (
	"context"
)

//定义服务接口
type SUrlsInf interface {
	Get(context.Context, string) (string, error)
	Set(context.Context, string) (string, error)
}

//定义中间件接口
type SUrlsMiddleware func(SUrlsInf) SUrlsInf
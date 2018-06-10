package svc

import (
	"time"
	"github.com/sirupsen/logrus"
	"context"
	"surls/global"
)

func LoggingMiddleware(logger global.Logger) SUrlsMiddleware{
	return func(next SUrlsInf) SUrlsInf {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger global.Logger
	Next SUrlsInf
}

func (mw loggingMiddleware) Set(ctx context.Context, s string) (entity SurlEntity, err error) {
	defer func(begin time.Time) {
		var info = make(logrus.Fields)
		info["method"] = "set"
		info["input"] = s
		info["output"] = entity
		info["err"] = err
		info["duration"] = time.Since(begin)
		global.Log.WithFields(info).Info("surls/v1/set")
	}(time.Now())

	entity, err = mw.Next.Set(ctx, s)
	return
}

func (mw loggingMiddleware) Get(ctx context.Context, s string) (entity SurlEntity, err error) {
	defer func(begin time.Time) {
		var info logrus.Fields = make(map[string]interface{})
		info["method"] = "get"
		info["input"] = s
		info["output"] = entity
		info["err"] = err
		info["duration"] = time.Since(begin)
		global.Log.WithFields(info).Info("surls/v1/get")
	}(time.Now())

	entity, err = mw.Next.Get(ctx, s)
	return
}


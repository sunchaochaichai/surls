package mw_svc

import (
	"time"
	"surls/svc/surlssvc/interfaces"
	"github.com/sirupsen/logrus"
	"context"
)

type LoggingMiddleware struct {
	Next interfaces.SUrlsInf
}

func (mw LoggingMiddleware) Set(ctx context.Context, s string) (output string, err error) {
	defer func(begin time.Time) {
		var info logrus.Fields = make(map[string]interface{})
		info["method"] = "set"
		info["input"] = s
		info["output"] = output
		info["err"] = err
		info["duration"] = time.Since(begin)
		logrus.WithFields(info).Info("surls/v1/set")
	}(time.Now())

	output, err = mw.Next.Set(ctx, s)
	return
}

func (mw LoggingMiddleware) Get(ctx context.Context, s string) (output string, err error) {
	defer func(begin time.Time) {
		var info logrus.Fields = make(map[string]interface{})
		info["method"] = "get"
		info["input"] = s
		info["output"] = output
		info["err"] = err
		info["duration"] = time.Since(begin)
		logrus.WithFields(info).Info("surls/v1/get")
	}(time.Now())

	output, err = mw.Next.Get(ctx, s)
	return
}

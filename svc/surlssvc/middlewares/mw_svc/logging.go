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
		logrus.Info(
			"method", "set",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Next.Set(ctx, s)
	return
}

func (mw LoggingMiddleware) Get(ctx context.Context, s string) (output string, err error) {
	defer func(begin time.Time) {
		logrus.Info(
			"method", "get",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Next.Get(ctx, s)
	return
}

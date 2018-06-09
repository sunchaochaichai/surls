package global

import (
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

var Redis *redis.Client

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors:    false,
		FullTimestamp:    true,
		TimestampFormat:  "2006-01-02 15:04:05.0000",
		QuoteEmptyFields: true,
	})

	LoadConf()

	Redis = NewRedisClient()
}

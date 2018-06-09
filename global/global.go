package global

import (
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

var Redis *redis.Client

var Logger *logrus.Logger

func init() {

	loadConf()

	Logger = newLogger()

	Redis = newRedisClient()
}

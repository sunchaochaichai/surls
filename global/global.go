package global

import (
	"github.com/go-redis/redis"
)

var Redis *redis.Client

var Log Logger

func init() {

	loadConf()

	Log = newLogger()

	Redis = newRedisClient()
}

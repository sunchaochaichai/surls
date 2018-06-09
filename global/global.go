package global

import (
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"os"
)

var Redis *redis.Client
var ProjectRealPath string

func init() {
	ProjectRealPath = os.Getenv("GOPATH") + "/src/surls"

	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors:    false,
		FullTimestamp:    true,
		TimestampFormat:  "2006-01-02 15:04:05.0000",
		QuoteEmptyFields: true,
	})

	LoadConf()

	Redis = NewRedisClient()
}

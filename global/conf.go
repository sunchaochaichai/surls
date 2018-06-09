package global

import (
	"surls/cli"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"log"
	"github.com/davecgh/go-spew/spew"
)

type conf struct {
	Redis redisConf
}

type redisConf struct {
	Addr     string
	Password string
	DB       int
}

var Conf conf

const (
	RUN_MODE_LOCAL     = "local"
	RUN_MODE_CONTAINER = "container"
)

func LoadConf() {
	log.Println("run mode:", cli.Params.RunMode)
	var confFile string
	switch cli.Params.RunMode {
	case RUN_MODE_LOCAL:
		confFile = "conf/local.yaml"
	case RUN_MODE_CONTAINER:
		confFile = "conf/container.yaml"
	default:
		log.Fatalln("unsuppoer run mode! use -h get help")
	}

	conf, _ := ioutil.ReadFile(confFile)
	err := yaml.Unmarshal(conf, &Conf)
	if err != nil {
		log.Fatalln("conf load failed", err)
	}

	spew.Dump(Conf)
}

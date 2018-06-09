package global

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/davecgh/go-spew/spew"
)

type conf struct {
	ProjectRealPath string
	Redis           redisConf
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

var ProjectRealPath = os.Getenv("GOPATH") + "/src/surls"
var RuntimeRealPath = ProjectRealPath + "/runtime"
var LogPath = RuntimeRealPath + "/logs"

func loadConf() {

	if err := godotenv.Load(ProjectRealPath + "/.env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	runMode := os.Getenv("RUN_MODE")
	log.Println("run mode:", runMode)

	var confFile string
	switch runMode {
	case RUN_MODE_LOCAL:
		confFile = ProjectRealPath + "/conf/local.yaml"
	case RUN_MODE_CONTAINER:
		confFile = ProjectRealPath + "/conf/container.yaml"
	default:
		log.Fatalln("unsuppoer run mode! supports:[local,container]")
	}

	conf, _ := ioutil.ReadFile(confFile)
	if err := yaml.Unmarshal(conf, &Conf); err != nil {
		log.Fatalln("conf load failed", err)
	}

	spew.Dump(Conf)
}

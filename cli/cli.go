package cli

import (
	"os"
	"github.com/urfave/cli"
	"log"
)

type params struct {
	GrpcServAddr string
	HttpServAddr string
	DebugAddr    string
	MetricsAddr  string
	MetricsPath  string
}

var Params params

func init() {

	app := cli.NewApp()
	app.Name = "SUrls"
	app.Usage = "go-kit 演示项目"
	app.Version = "v0.1.0"

	registFlags(app)

	app.Action = func(c *cli.Context) error {
		Params.GrpcServAddr = c.String("grpc-serv-addr")

		Params.HttpServAddr = c.String("http-serv-addr")

		Params.MetricsAddr = c.String("metrics-addr")

		Params.DebugAddr = c.String("debug-addr")

		Params.MetricsPath = c.String("metrics-path")

		return nil
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatalln("启动失败", err)
	}

	done()
}

func registFlags(app *cli.App) {
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "grpc-serv-addr",
			Usage: "grpc server addr",
			Value: ":7070",
		},
		cli.StringFlag{
			Name:  "http-serv-addr",
			Usage: "http server addr",
			Value: ":7071",
		},
		cli.StringFlag{
			Name:  "metrics-addr",
			Usage: "prometheus metrics addr",
			Value: ":7072",
		},
		cli.StringFlag{
			Name:  "debug-addr",
			Usage: "debug addr",
			Value: ":7073",
		},
		cli.StringFlag{
			Name:  "metrics-path",
			Usage: "prometheus metrics path",
			Value: "/metrics",
		},
	}
}

func done() {
	for _, v := range os.Args {
		if v == "-v" || v == "--version" {
			os.Exit(0)
		}

		if v == "-h" || v == "--help" {
			os.Exit(0)
		}
	}
}

package main

import (
	"log"
	"time"

	"github.com/asim/go-micro/v3"
	"github.com/micro/cli/v2"

	"github.com/jdxj/study_im/config"
	"github.com/jdxj/study_im/gate"
	"github.com/jdxj/study_im/logger"
)

const (
	modelName = "gate"
)

var (
	conf       *config.Config
	gateServer *gate.Gate
)

func main() {
	service := micro.NewService(
		micro.Flags(&cli.StringFlag{
			Name:  "configPath",
			Usage: "config path",
			Value: "./conf.yaml",
		}),

		micro.Name(modelName),

		micro.RegisterTTL(30*time.Second),
		micro.RegisterInterval(10*time.Second),

		// gate server
		micro.BeforeStart(func() error {
			gateCfg := conf.Gate
			gateServer = gate.New(gateCfg.Host, gateCfg.Port)
			gateServer.Run()
			return nil
		}),
		micro.AfterStop(func() error {
			gateServer.Stop()
			return nil
		}),

		// another
	)

	service.Init(
		micro.Action(func(ctx *cli.Context) error {
			path := ctx.String("configPath")
			c, err := config.New(path)
			if err != nil {
				return err
			}
			conf = c

			loggerCfg := conf.Logger
			logger.Init(loggerCfg.FileName, loggerCfg.AppName, loggerCfg.MaxSize,
				loggerCfg.MaxAge, loggerCfg.MaxBackups, loggerCfg.Level,
				loggerCfg.LocalTime, loggerCfg.Compress,
			)
			return nil
		}),
	)

	err := service.Run()
	if err != nil {
		log.Println(err)
	}
}

package main

import (
	"log"
	"time"

	"github.com/jdxj/study_im/proto/chat"

	"github.com/asim/go-micro/v3"
	"github.com/micro/cli/v2"

	"github.com/jdxj/study_im/config"
	"github.com/jdxj/study_im/logger"
)

const (
	modelName = "c2c"
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
	)

	service.Init(
		micro.Action(func(ctx *cli.Context) error {
			path := ctx.String("configPath")
			conf, err := config.New(path)
			if err != nil {
				return err
			}

			return Init(conf)
		}),
	)

	err := chat.RegisterC2CHandler(service.Server(), &C2CService{})
	if err != nil {
		log.Fatalln(err)
	}

	err = service.Run()
	if err != nil {
		log.Fatalln(err)
	}
}

func Init(conf *config.Config) error {
	loggerCfg := conf.Logger
	logger.Init(loggerCfg.FileName, loggerCfg.AppName, loggerCfg.MaxSize,
		loggerCfg.MaxAge, loggerCfg.MaxBackups, loggerCfg.Level,
		loggerCfg.LocalTime, loggerCfg.Compress,
	)

	return nil
}

package main

import (
	"log"
	"time"

	"github.com/asim/go-micro/v3"
	"github.com/micro/cli/v2"

	"github.com/jdxj/study_im/config"
	"github.com/jdxj/study_im/dao/mysql"
	"github.com/jdxj/study_im/dao/redis"
	"github.com/jdxj/study_im/logger"
	"github.com/jdxj/study_im/proto/chat"
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

	redisCfg := conf.Redis
	err := redis.Init(redisCfg.Pass, redisCfg.Host, redisCfg.Port, redisCfg.DB)
	if err != nil {
		return err
	}

	rabbitCfg := conf.Rabbit
	err = InitBroker(
		rabbitCfg.User, rabbitCfg.Pass, rabbitCfg.Host, "", rabbitCfg.Port)
	if err != nil {
		return err
	}

	mysqlCfg := conf.MySQL
	err = mysql.Init(
		mysqlCfg.User, mysqlCfg.Pass, mysqlCfg.Dbname, mysqlCfg.Host, mysqlCfg.Port)
	return err
}

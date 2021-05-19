package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jdxj/study_im/proto/chat"

	"github.com/jdxj/study_im/proto/login"

	"github.com/asim/go-micro/v3"
	"github.com/jdxj/study_im/dao/redis"
	"github.com/panjf2000/gnet"

	"github.com/jdxj/study_im/config"
	"github.com/jdxj/study_im/logger"
	"github.com/micro/cli/v2"
)

const (
	modelName = "gate"
)

var (
	conf *config.Config

	loginService login.LoginService
	c2cService   chat.C2CService
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

		micro.BeforeStart(func() error {
			gateCfg := conf.Gate
			gate, err := NewGate(gateCfg.Host, gateCfg.Port, gateCfg.Node)
			if err != nil {
				return err
			}
			err = broker.Subscribe(gate.handleBroker)
			if err != nil {
				return err
			}

			fmt.Printf("%#v\n", gateCfg)
			go func() {
				err := gate.Serve()
				if err != nil {
					log.Printf("start gate server failed: %s\n", err)
				}
			}()
			return nil
		}),
		micro.BeforeStop(func() error {
			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()

			addr := fmt.Sprintf("%s:%d", conf.Gate.Host, conf.Gate.Port)
			return gnet.Stop(ctx, addr)
		}),
	)

	service.Init(
		micro.Action(func(ctx *cli.Context) error {
			var err error
			path := ctx.String("configPath")
			conf, err = config.New(path)
			if err != nil {
				return err
			}

			return Init(conf)
		}),
	)

	loginService = login.NewLoginService("login", service.Client())
	c2cService = chat.NewC2CService("c2c", service.Client())

	err := service.Run()
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
	bindingKey := fmt.Sprintf("node.%d", conf.Gate.Node)
	err = InitBroker(
		rabbitCfg.User, rabbitCfg.Pass, rabbitCfg.Host, bindingKey, rabbitCfg.Port)
	return err
}

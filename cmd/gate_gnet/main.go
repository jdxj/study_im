package main

import (
	"log"

	"github.com/jdxj/study_im/config"
	"github.com/jdxj/study_im/logger"
)

func main() {
	// todo: 不要硬编码
	conf, err := config.New("./conf.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	loggerCfg := conf.Logger
	logger.Init(loggerCfg.FileName, loggerCfg.AppName, loggerCfg.MaxSize,
		loggerCfg.MaxAge, loggerCfg.MaxBackups, loggerCfg.Level,
		loggerCfg.LocalTime, loggerCfg.Compress,
	)

	gateCfg := conf.Gate
	gate, err := NewGate(gateCfg.Host, gateCfg.Port, gateCfg.Node)
	if err != nil {
		log.Fatalln(err)
	}

	err = gate.Serve()
	if err != nil {
		log.Fatalln(err)
	}
}

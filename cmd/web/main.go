package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/jdxj/study_im/cmd/web/api"
	"github.com/jdxj/study_im/config"
	"github.com/jdxj/study_im/logger"
)

func main() {
	conf, err := config.New("conf.yaml")
	if err != nil {
		log.Fatalf("New: %s", err)
	}

	loggerCfg := conf.Logger
	logger.Init(loggerCfg.FileName, loggerCfg.AppName, loggerCfg.MaxSize,
		loggerCfg.MaxAge, loggerCfg.MaxBackups, loggerCfg.Level,
		loggerCfg.LocalTime, loggerCfg.Compress,
	)

	webCfg := conf.Web
	gin.SetMode(webCfg.Mode)
	engine := api.NewServer()

	addr := fmt.Sprintf("%s:%d", webCfg.Host, webCfg.Port)
	err = engine.Run(addr)
	if err != nil {
		logger.Errorf("Run: %s", err)
	}
}

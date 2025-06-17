package main

import (
	"bin_blog/internal/config"
	"bin_blog/internal/pkg/log"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {

	// 加载配置
	err := config.LoadConfig()
	if err != nil {
		logrus.Fatal(err)
	}

	// 设置生产模式
	if config.GlobalConfig.Env == "prd" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.LoggerWithFormatter(log.logFormatter), gin.Recovery())
	err = router.Run(fmt.Sprintf(
		"%s:%s",
		config.GlobalConfig.Host,
		config.GlobalConfig.Port,
	))

	if err != nil {
		logrus.Fatal(err)
	}
	log.Logger.Info("用户登录成功", logrus.Fields{
		"user_id": 1001,
		"action":  "login",
	})
}

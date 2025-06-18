package main

import (
	"bin_blog/internal/config"
	"bin_blog/internal/middleware"
	"bin_blog/internal/pkg/dao"
	"bin_blog/internal/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	// 加载配置
	err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	logger.InitLogger()
	defer func(Log *zap.Logger) {
		err = Log.Sync()
		if err != nil {
			panic(err)
		}
	}(logger.Log)

	dao.InitMysql()

	// 设置生产模式
	if config.GlobalConfig.Env == "prd" {
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()
	}

	router := gin.New()
	router.Use(middleware.ZapLogger(), gin.Recovery())

	err = router.Run(fmt.Sprintf(
		"%s:%s",
		config.GlobalConfig.Host,
		config.GlobalConfig.Port,
	))

	if err != nil {
		panic(err)
	}
}

package main

import (
	"bin_blog/internal/config"
	"bin_blog/internal/pkg/dao"
	"bin_blog/internal/pkg/logger"
	"bin_blog/internal/routers"
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

	// 初始化 zap
	logger.InitLogger()
	defer func(Log *zap.Logger) {
		err = Log.Sync()
		if err != nil {
			panic(err)
		}
	}(logger.Log)

	// 初始化 Mysql
	dao.InitMysql()

	// 设置生产模式
	if config.GlobalConfig.Env == "prd" {
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()
	}

	router := gin.New()
	router.Use(logger.ZapLogger(), gin.Recovery())

	// 初始化路由
	routers.InitRouter(router)

	err = router.Run(fmt.Sprintf(
		"%s:%s",
		config.GlobalConfig.Host,
		config.GlobalConfig.Port,
	))

	if err != nil {
		panic(err)
	}
}

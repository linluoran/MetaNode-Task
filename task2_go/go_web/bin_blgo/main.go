package main

import (
	"bin_blog/internal/config"
	"bin_blog/internal/pkg/dao"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zeromicro/go-zero/core/logx"
)

func main() {

	logx.MustSetup(logx.LogConf{
		Mode:     "file",   // 文件模式（非控制台）
		Path:     "./logs", // 日志目录
		Level:    "info",   // 记录 Info 及以上级别
		Encoding: "plain",  // 纯文本格式（无 JSON 冗余）
	})

	// 加载配置
	err := config.LoadConfig()
	if err != nil {
		logx.Error(err)
	}
	dao.InitMysql()

	// 设置生产模式
	if config.GlobalConfig.Env == "prd" {
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()
	}

	router := gin.New()
	router.Use(gin.Recovery())

	err = router.Run(fmt.Sprintf(
		"%s:%s",
		config.GlobalConfig.Host,
		config.GlobalConfig.Port,
	))

	if err != nil {
		logx.Error(err)
	}
	logx.Error("用户登录成功", "id=1001")
}

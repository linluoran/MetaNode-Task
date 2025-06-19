package logger

import (
	"gin_blog/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"path"
	"time"
)

var Log *zap.Logger

func InitLogger() {
	// 确保logs目录存在
	logConf := config.GlobalConfig.Log

	if err := os.MkdirAll(logConf.LogPath, 0755); err != nil {
		log.Fatalf("创建日志目录失败: %v", err)
	}

	// 生成带日期的日志文件名
	currentDate := time.Now().Format("20060102")
	logFileName := path.Join(logConf.LogPath, "gin"+currentDate+".log")
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logFileName,
		MaxSize:    logConf.MaxSize,
		MaxBackups: logConf.MaxBackups,
		MaxAge:     logConf.MaxAge,
		Compress:   logConf.Compress,
	}

	// 自定义 Zap 配置
	cfg := zap.NewProductionConfig()
	if config.GlobalConfig.Env == "dev" {
		// 开发环境输出 Debug 日志
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	cfg.OutputPaths = []string{"stdout", lumberJackLogger.Filename} // 同时输出到控制台和文件
	Log, _ = cfg.Build()

	// 替换全局 logger
	zap.ReplaceGlobals(Log)
}

func ZapLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		p := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		// 处理请求
		c.Next()

		duration := time.Since(start)

		// 记录日志
		zap.L().Info("HTTP Request",
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", p),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.Duration("duration", duration),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
		)
	}
}

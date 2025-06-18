package logger

import (
	"bin_blog/internal/config"
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

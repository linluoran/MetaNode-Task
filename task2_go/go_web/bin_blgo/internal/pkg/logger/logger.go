package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"path"
	"time"
)

var LogFile *os.File

func InitLog() {
	// 确保logs目录存在
	if err := os.MkdirAll("logs", 0755); err != nil {
		log.Fatalf("创建日志目录失败: %v", err)
	}

	// 生成带日期的日志文件名
	currentDate := time.Now().Format("20060102")
	logFileName := path.Join("logs", "gin"+currentDate+".log")

	// 打开日志文件
	logFile, err := os.OpenFile(
		logFileName,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)
	if err != nil {
		log.Fatalf("打开日志文件失败: %v", err)
	}

	LogFile = logFile
	// 设置日志输出
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)
}

func Formatter(param gin.LogFormatterParams) string {
	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
		param.ClientIP,
		param.TimeStamp.Format(time.RFC1123),
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		param.Request.UserAgent(),
		param.ErrorMessage,
	)
}

func CloseLogger() {
	if LogFile != nil {
		_ = LogFile.Close()
	}
}

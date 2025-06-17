package log

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
	"time"
)

var Logger = logrus.New()

func init() {
	// 创建logs目录
	logDir := "logs"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		logrus.Fatal("创建日志目录失败: ", err)
	}

	// 打开日志文件
	logFile, err := os.OpenFile(
		filepath.Join(logDir, "app.log"),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666,
	)
	if err != nil {
		logrus.Fatal("打开日志文件失败: ", err)
	}
	defer logFile.Close()

	// 配置JSON格式
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05", // 自定义时间戳
	})
	logrus.SetOutput(io.MultiWriter(logFile, os.Stdout))
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetReportCaller(true)
	logrus.Info("nimabi")

}
func logFormatter(param gin.LogFormatterParams) string {
	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
		param.ClientIP,
		param.TimeStamp.Format(time.RFC3339),
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		param.Request.UserAgent(),
		param.ErrorMessage,
	)
}

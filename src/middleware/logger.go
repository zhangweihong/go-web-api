package middleware

import (
	"gin-framework/basic/src/helper"
	"io"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	lfshook "github.com/rifflock/lfshook"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func LoggerToFile() gin.HandlerFunc {
	logFilePath := helper.GetLogFilePath()
	fileName := logFilePath
	// 实例化
	Logger = logrus.New()

	// 设置输出
	Logger.SetOutput(io.MultiWriter(os.Stdout))

	Logger.Formatter = &logrus.TextFormatter{
		ForceColors: true,
	}

	// 设置日志级别
	Logger.SetLevel(logrus.InfoLevel)

	// 设置 rotatelogs
	errorLogWriter, _ := rotatelogs.New(
		// 分割后的文件名称
		fileName+"_error.%Y%m%d.log",

		// // 生成软链，指向最新日志文件
		// rotatelogs.WithLinkName(fileName),

		// 设置最大保存时间(60天)
		rotatelogs.WithMaxAge(60*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	accessLogWriter, _ := rotatelogs.New(
		// 分割后的文件名称
		fileName+"_access.%Y%m%d.log",

		// // 生成软链，指向最新日志文件
		// rotatelogs.WithLinkName(fileName),

		// 设置最大保存时间(60天)
		rotatelogs.WithMaxAge(60*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  accessLogWriter,
		logrus.FatalLevel: errorLogWriter,
		logrus.DebugLevel: errorLogWriter,
		logrus.WarnLevel:  errorLogWriter,
		logrus.ErrorLevel: errorLogWriter,
		logrus.PanicLevel: errorLogWriter,
	}

	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// 新增 Hook
	Logger.AddHook(lfHook)

	return func(c *gin.Context) {
		// 处理请求
		c.Next()
		go func() {
			// 开始时间
			startTime := time.Now()
			// 结束时间
			endTime := time.Now()

			// 执行时间
			latencyTime := endTime.Sub(startTime)

			// 请求方式
			reqMethod := c.Request.Method

			// 请求路由
			reqUri := c.Request.RequestURI

			// 状态码
			statusCode := c.Writer.Status()

			// 请求IP
			clientIP := c.ClientIP()
			// 日志格式
			Logger.WithFields(logrus.Fields{
				"status_code":  statusCode,
				"latency_time": latencyTime,
				"client_ip":    clientIP,
				"req_method":   reqMethod,
				"req_uri":      reqUri,
			}).Info()
		}()
	}
}

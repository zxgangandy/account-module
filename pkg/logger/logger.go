package logger

import (
	"account-module/pkg/conf"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"path"
	"time"
)

// 日志配置
func Init() {
	config := conf.GetConfig()

	logrus.SetReportCaller(true)

	var formatter logrus.Formatter
	//var stdFormatter *prefixed.TextFormatter
	//var formatter *prefixed.TextFormatter

	switch config.Application.LogOutPath {
	case "file":
		formatter = &logrus.JSONFormatter{}
	default:
		formatter = &logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02.15:04:05.000000",
			ForceColors:     true,
			DisableColors:   false,
		}
	}

	logrus.SetFormatter(formatter)

	switch config.Application.LogLevel {
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "panic":
		logrus.SetLevel(logrus.PanicLevel)
	case "fatal":
		logrus.SetLevel(logrus.FatalLevel)
	default:
		logrus.SetLevel(logrus.DebugLevel)
	}

	switch config.Application.LogOutPath {
	case "file":
		logFileName := path.Join("logs", config.Application.Name+".log")
		var logMaxSaveDay = time.Duration(config.Application.LogMaxSaveDay)
		logWriter, err := rotatelogs.New(
			logFileName+".%Y-%m-%d",                                  // 日志切割名称
			rotatelogs.WithLinkName(logFileName),                     // 生成软链，指向最新日志文件
			rotatelogs.WithMaxAge(logMaxSaveDay*24*time.Hour),        // 文件最大保存时间
			rotatelogs.WithRotationTime(time.Duration(24)*time.Hour), // 日志切割时间间隔
		)
		if err != nil {
			log.Fatal("Create rotate logs object fail: ", err)
		}

		//	logrus.SetOutput(logWriter)

		writeMap := lfshook.WriterMap{
			logrus.InfoLevel:  logWriter,
			logrus.FatalLevel: logWriter,
			logrus.DebugLevel: logWriter,
			logrus.WarnLevel:  logWriter,
			logrus.ErrorLevel: logWriter,
			logrus.PanicLevel: logWriter,
		}

		lfHook := lfshook.NewHook(writeMap, formatter)
		logrus.AddHook(lfHook)

	default:
		logrus.SetOutput(os.Stdout)
	}
}

// @File:  logger.go
// @Time:  2022/1/5 12:18 PM
// @Author: ClassmateLin
// @Email: classmatelin.site@gmail.com
// @Site: https://www.classmatelin.top
// @Description:
// @Cron: * */1 * * *

package global

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"scripts/config"
	"time"
)

var Log *logrus.Logger

// init
// @description: 初始化日志
func init() {
	logDirectory := config.VP.Get(`logger.directory`)
	logFilename := config.VP.Get(`logger.filename`)

	// 日志文件
	fileName := path.Join(logDirectory.(string), logFilename.(string))

	// 写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	// 实例化
	Log = logrus.New()

	// 设置输出
	Log.Out = src

	// 设置日志级别
	Log.SetLevel(logrus.DebugLevel)

	// 设置 rotatelogs
	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		fileName+".%Y%m%d.log",

		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),

		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	lfHook := lfshook.NewHook(
		writeMap,
		&logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		},
	)
	// 新增 Hook
	Log.Hooks.Add(lfHook)
}

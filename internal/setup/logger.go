package setup

import (
	"gin-admin/global"
	"gin-admin/pkg/logger"
	"io"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	w := &lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   500,
		MaxAge:    10,
		LocalTime: true,
	}
	opts := &logger.Options{
		Level:             global.LoggerSetting.Level,
		DisableCaller:     global.LoggerSetting.DisableCaller,
		DisableStacktrace: global.LoggerSetting.DisableStacktrace,
		EnableColor:       global.LoggerSetting.EnableColor,
		Development:       global.LoggerSetting.Development,
		WriteSyncer:       []io.Writer{os.Stdout, w},
	}
	global.Logger = logger.New(opts)

	return nil
}

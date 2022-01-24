package app

import (
	"fmt"
	"github.com/gin-pro/gin-pro-base/core/logx"
	"github.com/sirupsen/logrus"
	"path/filepath"
)

func InitLog() error {
	dir := "logs"
	mp := logx.PathMap{
		logrus.InfoLevel:  filepath.Join(dir, "info.log"),
		logrus.ErrorLevel: filepath.Join(dir, "error.log"),
		logrus.DebugLevel: filepath.Join(dir, "debug.log"),
		logrus.WarnLevel:  filepath.Join(dir, "warn.log"),
	}
	logger, err := logx.InitLogWithFile(mp)
	if err != nil {
		fmt.Println("init logs err", err)
		return err
	}
	Log = logger
	return nil
}

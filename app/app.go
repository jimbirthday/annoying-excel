package app

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"xorm.io/xorm"
)

var (
	Cnf conf
	Ctx context.Context
	Cnl context.CancelFunc

	Log *logrus.Logger
	DB  *xorm.Engine
)

func Init() {
	Ctx, Cnl = context.WithCancel(context.Background())
	err := initYaml()
	if err != nil {
		fmt.Println("initYaml err:", err.Error())
		return
	}
}

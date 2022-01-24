package app

import (
	"fmt"
	"github.com/gin-pro/gin-pro-base/core/stores/sqlx"
)

func InitDB() error {
	cnf := &sqlx.MysqlConf{
		Host:     Cnf.Database.Host,
		Port:     Cnf.Database.Port,
		Database: Cnf.Database.Database,
		Username: Cnf.Database.Username,
		Password: Cnf.Database.Password,
		ShowSQL:  Cnf.Database.ShowSQL,
	}
	db, err := sqlx.InitMysql(cnf)
	if err != nil {
		fmt.Println("init db err", err)
		return err
	}
	db.ShowSQL(cnf.ShowSQL)
	DB = db
	return nil
}

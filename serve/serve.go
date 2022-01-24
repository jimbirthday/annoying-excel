package serve

import (
	"fmt"
	"gridtransaction/app"
	"gridtransaction/base"
	"gridtransaction/routers"
)

func Run() {
	app.Init()
	err := app.InitDB()
	if err != nil {
		fmt.Println("db err:", err.Error())
		return
	}
	err = app.InitLog()
	if err != nil {
		fmt.Println("db err:", err.Error())
		return
	}
	err = base.InitHttpClient()
	if err != nil {
		fmt.Println("InitHttpClient err:", err.Error())
		return
	}

	eg := routers.InitRouter()
	eg.Run(fmt.Sprintf(":%d", app.Cnf.Server.Port))
}

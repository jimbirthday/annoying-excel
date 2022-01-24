package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-pro/gin-pro-base/routerx"
	"gridtransaction/routers/user"
)

func InitRouter() *routerx.ApiEngine {
	engine := routerx.Default()
	r := engine.NewGroup("/api")
	testc(r)
	r.Use(MidAccessAllowFun)
	basec(r)
	userc(r)
	return engine
}

func testc(r *routerx.ApiGroup) {
	r.GET("test", func(c *gin.Context) {
		c.JSON(200, "test")
	})
	r.GET("ping", func(c *gin.Context) {
		c.JSON(200, "pong")
	})
}

func basec(r *routerx.ApiGroup) {
	r.POST("login", user.Login)
	r.POST("register", user.Register)
	r.Any("isLogin", IsLogin)
}

func userc(r *routerx.ApiGroup) {
	group := r.NewGroup("/user")
	group.Use(MidTokenFun)
}

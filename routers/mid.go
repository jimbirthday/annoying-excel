package routers

import (
	"github.com/gin-gonic/gin"
	"gridtransaction/comm"
	"gridtransaction/utils"
	"net/http"
	"strings"
)

func MidAccessAllowFun(c *gin.Context) {
	method := strings.ToUpper(c.Request.Method)
	if method == "OPTIONS" || method == "POST" {
		c.Header("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
		c.Header("Access-Control-Allow-Headers", "*,Content-Type")
		c.Header("Access-Control-Allow-Methods", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
	}
	//放行所有OPTIONS方法
	if method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}
	// 处理请求
	c.Next()
}

func MidTokenFun(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		comm.NewResponse(comm.ResponseCodeAuth).ResJson(c)
		c.Abort()
		return
	}
	if !utils.CompareToken(auth) {
		comm.NewResponse(comm.ResponseCodeAuth).ResJson(c)
		c.Abort()
		return
	}
	// 处理请求
	c.Next()
}

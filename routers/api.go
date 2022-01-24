package routers

import (
	"github.com/gin-gonic/gin"
	"gridtransaction/comm"
	"gridtransaction/utils"
)

func IsLogin(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		comm.NewResponseOk(false).ResJson(c)
		return
	}
	comm.NewResponseOk(utils.CompareToken(auth)).ResJson(c)
	return
}

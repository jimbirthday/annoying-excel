package user

import (
	"github.com/gin-gonic/gin"
	"gridtransaction/app"
	"gridtransaction/bean"
	"gridtransaction/comm"
	"gridtransaction/model"
	"gridtransaction/models"
	"gridtransaction/utils"
	"time"
)

func Register(c *gin.Context, p *bean.RegisterParam) {
	err := p.Check()
	if err != nil {
		comm.NewResponse(comm.ResponseCodeErr).
			SetMsgf(err.Error()).
			ResJson(c)
		return
	}
	tu := &model.TUser{}
	_, err = app.DB.Where("user_name = ?", p.UserName).Get(tu)
	if err != nil {
		comm.NewResponse(comm.ResponseCodeErr).SetMsgf(err.Error()).ResJson(c)
		return
	}
	if tu.UserName == p.UserName {
		comm.NewResponse(comm.ResponseCodeErr).SetMsgf("用户名重复").ResJson(c)
		return
	}
	tu.UserName = p.UserName
	tu.Nick = p.Nick
	if p.Nick == "" {
		p.Nick = utils.GenerateNick()
	}

	password, err := utils.GenerateFromPassword(p.Password)
	if err != nil {
		comm.NewResponse(comm.ResponseCodeErr).SetMsgf(err.Error()).ResJson(c)
		return
	}

	tu.Password = password
	tu.CreateTime = time.Now()
	_, err = app.DB.Insert(tu)
	if err != nil {
		comm.NewResponse(comm.ResponseCodeErr).SetMsgf(err.Error()).ResJson(c)
		return
	}
	comm.NewResponseOk().ResJson(c)

}

func Login(c *gin.Context, p *bean.LoginParam) {
	err := p.Check()
	if err != nil {
		comm.NewResponse(comm.ResponseCodeErr).
			SetMsgf(err.Error()).
			ResJson(c)
		return
	}
	tu := &model.TUser{
		UserName: p.UserName,
	}
	get, err := app.DB.Get(tu)
	if err != nil {
		comm.NewResponse(comm.ResponseCodeErr).
			SetMsgf(err.Error()).
			ResJson(c)
		return
	}
	if !get {
		comm.NewResponse(comm.ResponseCodeErrs).SetMsgf("账号不正确").
			ResJson(c)
		return
	}
	if !utils.CompareHashAndPassword(tu.Password, p.Password) {
		comm.NewResponse(comm.ResponseCodeErr).
			SetMsgf("密码不正确").
			ResJson(c)
		return
	}
	token, err := utils.GenerateToken(tu)
	if err != nil {
		comm.NewResponse(comm.ResponseCodeErr).
			SetMsgf(err.Error()).
			ResJson(c)
		return
	}
	info := &models.LoginInfo{
		Token: token,
	}
	location, _ := time.LoadLocation("Asia/Shanghai")
	tu.LastLogin = time.Now().In(location)
	app.DB.Where("id = ?", tu.Id).Cols("last_login").Update(tu)
	comm.NewResponseOk(info).ResJson(c)
}

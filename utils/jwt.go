package utils

import (
	"encoding/json"
	"errors"
	ginpro "github.com/gin-pro/gin-pro-base"
	"gridtransaction/model"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gin-pro/gin-pro-base/core/jwtx"
)

const TokenKey = "ginpro"

func GenerateToken(tu *model.TUser) (string, error) {
	m := jwt.MapClaims{}
	u := map[string]interface{}{}
	u["id"] = tu.Id
	u["nick"] = tu.Nick
	u["lastLogin"] = tu.LastLogin
	u["userName"] = tu.UserName
	m["user"] = u
	token, err := jwtx.CreateToken(m, TokenKey, time.Hour*6)
	if err != nil {
		return "", err
	}
	return token, nil
}

func CompareToken(tk string) bool {
	m := jwtx.GetToken(tk, TokenKey)
	if m == nil {
		return false
	}
	t, ok := m["timeout"]
	if !ok {
		return false
	}
	ti, ok := t.(string)
	if !ok {
		return false
	}
	tms, _ := time.Parse(time.RFC3339Nano, ti)
	return tms.After(time.Now())
}

func GetCurUser(c *gin.Context) (*model.TUser, error) {
	defer ginpro.Recover()
	tk := c.GetHeader("Authorization")
	m := jwtx.GetToken(tk, TokenKey)
	if m == nil {
		return nil, errors.New("GetToken si empty")
	}
	t, ok := m["user"]
	if !ok {
		return nil, errors.New("user is empty")
	}
	tu := &model.TUser{}
	marshal, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(marshal, tu)
	if err != nil {
		return nil, err
	}
	return tu, nil
}

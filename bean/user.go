package bean

import (
	"fmt"
	"gridtransaction/comm"
	"regexp"
)

type RegisterParam struct {
	UserName   string `json:"userName"`
	Password   string `json:"password"`
	Nick       string `json:"nick"`
	RePassword string `json:"rePassword"`
}

func (p *RegisterParam) Check() error {
	err := checkUserName(p.UserName)
	if err != nil {
		return err
	}
	if p.Password != p.RePassword {
		return fmt.Errorf("两次密码不一致")
	}
	err = checkUserName(p.Password)
	if err != nil {
		return err
	}
	err = checkUserName(p.RePassword)
	if err != nil {
		return err
	}
	err = checkNick(p.Nick)
	if err != nil {
		return err
	}
	return nil
}

type LoginParam struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

func (p *LoginParam) Check() error {
	err := checkUserName(p.UserName)
	if err != nil {
		return err
	}
	err = checkUserName(p.Password)
	if err != nil {
		return err
	}
	return nil
}

func checkNick(nick string) error {
	if nick == "" {
		return nil
	}
	compile := regexp.MustCompile(comm.RegNick)
	if !compile.MatchString(nick) {
		return fmt.Errorf("昵称长度3~20之间")
	}
	return nil
}

func checkUserName(userName string) error {
	if userName == "" {
		return fmt.Errorf("账号不能为空")
	}
	compile := regexp.MustCompile(comm.RegUserName)
	if !compile.MatchString(userName) {
		return fmt.Errorf("账号字母开头，长度5~16之间，允许字母数字下划线")
	}
	return nil
}

func checkPassword(password string) error {
	compile := regexp.MustCompile(comm.RegPassword)
	if password == "" {
		return fmt.Errorf("密码不能为空")
	}
	if !compile.MatchString(password) {
		return fmt.Errorf("密码字母开头，允许6~18之间，允许字母数字下划线")
	}
	return nil
}

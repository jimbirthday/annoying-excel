package model

import (
	"time"
)

type TUser struct {
	Id         int       `xorm:"not null pk autoincr INT(11)" json:"id"`
	UserName   string    `xorm:"VARCHAR(255)" json:"userName"`
	Password   string    `xorm:"VARCHAR(255)" json:"password"`
	Nick       string    `xorm:"VARCHAR(255)" json:"nick"`
	Avatar     string    `xorm:"VARCHAR(255)" json:"avatar"`
	LastLogin  time.Time `xorm:"DATETIME" json:"lastLogin"`
	CreateTime time.Time `xorm:"DATETIME" json:"createTime"`
}

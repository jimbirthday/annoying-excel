package comm

import (
	"regexp"
	"time"
)

const (
	TimeZero   = -62135596800
	TimeFmt    = "2006-01-02 15:04:05"
	TimeFmts   = "2006-01-02"
	TimeFmtm   = "2006-01"
	TimeFmtt   = "20060102150405"
	TimeFmtpck = "2006-01-02T15:04:05.999999999Z"
)

var (
	TimeZHT  = time.FixedZone("CST", 8*3600)
	RegPhone = regexp.MustCompile(`^([0-9]{0,3})[1][3,4,5,7,8,9][0-9]{9}$`)
)

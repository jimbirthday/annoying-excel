package utils

import (
	"gridtransaction/comm"
	"time"
)

func GetNowDayStart() string {
	return GetDayStart(time.Now())
}

func GetNowDayEnd() string {
	return GetDayEnd(time.Now())
}

func GetDayStart(t time.Time) string {
	startTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Format(comm.TimeFmt)
	return startTime
}
func GetDayEnd(t time.Time) string {
	endTime := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location()).Format(comm.TimeFmt)
	return endTime
}

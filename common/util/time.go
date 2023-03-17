package util

import (
	"time"
)

var timeTemplate1 = "2006-01-02 15:04:05"
var timeTemplate2 = "2006/01/02 15:04:05"
var timeTemplate3 = "2006-01-02 "
var timeTemplate4 = "20060102150405"
var timeTemplate5 = "15"
var timeTemplate6 = "2006-01-02"
var timeTemplate7 = "15:04:05"
var timeTemplate8 = "20060102"

func TimeToInt64(tm time.Time) int64 {
	return tm.Unix()
}

func Int64ToTime(in int64) time.Time {
	return time.Unix(in, 0)
}

func TimeToStringTemplate1(tm time.Time) string {
	return tm.Format(timeTemplate1)
}
func TimeToStringTemplate2(tm time.Time) string {
	return tm.Format(timeTemplate2)
}
func TimeToStringTemplate3(tm time.Time) string {
	return tm.Format(timeTemplate3)
}
func TimeToStringTemplate4(tm time.Time) string {
	return tm.Format(timeTemplate4)
}
func TimeToStringTemplate5(tm time.Time) string {
	return tm.Format(timeTemplate5)
}
func TimeToStringTemplate6(tm time.Time) string {
	return tm.Format(timeTemplate6)
}
func TimeToStringTemplate7(tm time.Time) string {
	return tm.Format(timeTemplate7)
}
func TimeToStringTemplate8(tm time.Time) string {
	return tm.Format(timeTemplate8)
}
func StringTemplate1ToTime(tmStr string) (tm time.Time, err error) {
	tm, err = time.ParseInLocation(timeTemplate1, tmStr, time.Local)
	return
}
func StringTemplate2ToTime(tmStr string) (tm time.Time, err error) {
	tm, err = time.ParseInLocation(timeTemplate2, tmStr, time.Local)
	return
}
func StringTemplate3ToTime(tmStr string) (tm time.Time, err error) {
	tm, err = time.ParseInLocation(timeTemplate3, tmStr, time.Local)
	return
}

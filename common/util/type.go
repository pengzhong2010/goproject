package util

import (
	"fmt"
	"strconv"
)

func StringToInt64(in string) (out int64, err error) {
	out, err = strconv.ParseInt(in, 10, 64)
	return
}

func StringToUint64(in string) (out uint64, err error) {
	out, err = strconv.ParseUint(in, 10, 64)
	return
}
func StringToInt(in string) (out int, err error) {
	out, err = strconv.Atoi(in)
	return
}
func Int64ToString(in int64) (out string) {
	out = strconv.FormatInt(in, 10)
	return
}
func Uint64ToString(in uint64) (out string) {
	out = strconv.FormatUint(in, 10)
	return
}
func IntToString(in int) (out string) {
	out = strconv.Itoa(in)
	return
}
func Float64ToString(in float64) (out string) {
	out = strconv.FormatFloat(in, 'f', -1, 32)
	return
}
func StringToFloat64(in string) (out float64, err error) {
	out, err = strconv.ParseFloat(in, 64)
	return
}
func InterfaceNumToInt64(in interface{}) (out int64, err error) {
	var (
		f  float64
		ok bool
	)
	if f, ok = in.(float64); !ok {
		err = fmt.Errorf("float64 judge err")
		return
	}
	out, err = StringToInt64(Float64ToString(f))
	return
}
func InterfaceNumToString(in interface{}) (out string, err error) {
	var (
		f  float64
		ok bool
	)
	if f, ok = in.(float64); !ok {
		err = fmt.Errorf("float64 judge err")
		return
	}
	out = Float64ToString(f)
	return
}

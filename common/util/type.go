package util

import "strconv"

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

package util

import "os"

func GetEnv(key, defVal string) (value string) {
	var ok bool
	if value, ok = os.LookupEnv(key); !ok {
		value = defVal
	}
	return
}

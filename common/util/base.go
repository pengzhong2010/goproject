package util

import (
	"encoding/json"
	"reflect"
)

func GetIndex(xs interface{}, x interface{}) int64 {
	if sl := reflect.ValueOf(xs); sl.Kind() == reflect.Slice {
		for i := 0; i < sl.Len(); i++ {
			switch y := sl.Index(i).Interface().(type) {
			case int:
				if y == x.(int) {
					return int64(i)
				}
			case string:
				if y == x.(string) {
					return int64(i)
				}
			case int64:
				if y == x.(int64) {
					return int64(i)
				}
			}
		}
	}
	return -1
}
func StructTypeTrans(in interface{}, out interface{}) (err error) {
	b, err := json.Marshal(in)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, out)
	if err != nil {
		return
	}
	return
}

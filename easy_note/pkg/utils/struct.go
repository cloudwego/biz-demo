package utils

import "github.com/bytedance/sonic"

func StructToString(s interface{}) string {
	bytes, _ := sonic.Marshal(s)
	return string(bytes)
}

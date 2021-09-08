package utils

import (
	"fmt"
	"runtime"
)

func GetFuncName() string {
	pt, _, _, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("fail to get stack trace")
		return ""
	}

	funcName := runtime.FuncForPC(pt).Name()

	return funcName
}

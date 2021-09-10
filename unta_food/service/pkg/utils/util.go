package utils

import (
	"fmt"
	"log"
	"net/url"
	"runtime"
	"strings"
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

func SplitMultiSep(s string, sep []string) []string {
	log.Printf("[START] :%s", GetFuncName())
	defer log.Printf("[END] :%s", GetFuncName())

	var ret []string
	ret = strings.Split(s, sep[0])
	if len(sep) > 1 {
		ret2 := []string{}
		for _, r := range ret {
			ret2 = append(ret2, SplitMultiSep(r, sep[1:])...)
		}
		ret = ret2
	}
	return ret
}

func IsURL(s string) bool {
	log.Printf("[START] :%s", GetFuncName())
	defer log.Printf("[END] :%s", GetFuncName())

	u, err := url.Parse(s)
	return err == nil && u.Scheme != "" && u.Host != ""
}

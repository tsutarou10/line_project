package controller

import (
	"log"
	"net/url"

	"github.com/tsutarou10/line_project/service/pkg/utils"
)

func isURL(s string) bool {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	u, err := url.Parse(s)
	return err == nil && u.Scheme != "" && u.Host != ""
}

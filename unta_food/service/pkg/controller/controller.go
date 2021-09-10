package controller

import (
	"log"

	"github.com/tsutarou10/line_project/service/pkg/usecase"
	"github.com/tsutarou10/line_project/service/pkg/utils"
)

type Controller struct {
	in usecase.UTNAFoodInputPort
}

func NewController(in usecase.UTNAFoodInputPort) *Controller {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	return &Controller{in}
}

func createMemo(src []string) string {
	rsl := ""
	for i, s := range src {
		rsl += s
		if i != len(src)-1 {
			rsl += " "
		}
	}
	return rsl
}

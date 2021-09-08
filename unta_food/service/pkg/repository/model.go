package repository

import (
	"log"

	"github.com/tsutarou10/line_project/service/pkg/entity"
	"github.com/tsutarou10/line_project/service/pkg/utils"
)

func toModel(input entity.RegisterEntity) testInput {
	log.Printf("[START] :%s", utils.GetFuncName())
	//defer log.Printf("[END] :%s", utils.GetFuncName())

	return testInput{
		URL: input.URL,
	}
}

package repository

import (
	"log"
	"time"

	"github.com/tsutarou10/line_project/service/pkg/entity"
	"github.com/tsutarou10/line_project/service/pkg/utils"
)

type utnaFoodSchema struct {
	ID        int64  `dynamo:"id"`
	URL       string `dynamo:"url" index:"URLIndex"`
	Memo      string `dynamo:"memo"`
	UpdatedAt int64  `dynamo:"updatedAt"`
}

type utnaFoodRegisterStatus struct {
	Status    string `dynamo:"status"`
	Number    int64  `dynamo:"number"`
	UpdatedAt int64  `dynamo:"updatedAt"`
}

func toModel(input entity.UTNAEntityFood, id int64) utnaFoodSchema {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	return utnaFoodSchema{
		ID:        id,
		URL:       input.URL,
		Memo:      input.Memo,
		UpdatedAt: time.Now().Unix(),
	}
}

func toEntity(input utnaFoodSchema) entity.UTNAEntityFood {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	return entity.UTNAEntityFood{
		ID:   input.ID,
		URL:  input.URL,
		Memo: input.Memo,
	}
}

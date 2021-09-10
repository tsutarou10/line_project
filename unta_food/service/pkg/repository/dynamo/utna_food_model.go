package dynamo

import (
	"log"

	"github.com/tsutarou10/line_project/service/pkg/entity"
	"github.com/tsutarou10/line_project/service/pkg/utils"
)

type utnaFood struct {
	URL         string `dynamo:"url"`
	ImageURL    string `dynamo:"imageUrl"`
	Title       string `dynamo:"title"`
	Memo        string `dynamo:"memo"`
	IsCompleted bool   `dynamo:"isCompleted"`
	UpdatedAt   int64  `dynamo:"updatedAt"`
}

func toModel(input entity.UTNAEntityFood) utnaFood {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	return utnaFood{
		URL:         input.URL,
		ImageURL:    input.ImageURL,
		Title:       input.Title,
		IsCompleted: input.IsCompleted,
		Memo:        input.Memo,
		UpdatedAt:   input.UpdatedAt,
	}
}

func toEntity(input utnaFood) entity.UTNAEntityFood {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	return entity.UTNAEntityFood{
		URL:         input.URL,
		ImageURL:    input.ImageURL,
		Title:       input.Title,
		IsCompleted: input.IsCompleted,
		Memo:        input.Memo,
		UpdatedAt:   input.UpdatedAt,
	}
}

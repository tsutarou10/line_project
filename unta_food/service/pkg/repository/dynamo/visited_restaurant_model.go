package dynamo

import (
	"log"

	"github.com/tsutarou10/line_project/service/pkg/entity"
	"github.com/tsutarou10/line_project/service/pkg/utils"
)

type visitedRestaurant struct {
	URL       string `dynamo:"url"`
	ImageURL  string `dynamo:"imageUrl"`
	Title     string `dynamo:"title"`
	Memo      string `dynamo:"memo"`
	UpdatedAt int64  `dynamo:"updatedAt"`
}

func toModelOfVisitedRestaurant(input entity.UTNAEntityFood) visitedRestaurant {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	return visitedRestaurant{
		URL:       input.URL,
		ImageURL:  input.ImageURL,
		Title:     input.Title,
		Memo:      input.Memo,
		UpdatedAt: input.UpdatedAt,
	}
}

func toEntity(input visitedRestaurant) entity.UTNAEntityFood {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	return entity.UTNAEntityFood{
		URL:       input.URL,
		ImageURL:  input.ImageURL,
		Title:     input.Title,
		Memo:      input.Memo,
		UpdatedAt: input.UpdatedAt,
	}
}

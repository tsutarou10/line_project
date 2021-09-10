package dynamo

import (
	"log"
	"time"

	"github.com/tsutarou10/line_project/service/pkg/entity"
	"github.com/tsutarou10/line_project/service/pkg/utils"
)

type utnaFood struct {
	ID        int64  `dynamo:"id"`
	URL       string `dynamo:"url" index:"URLIndex"`
	ImageURL  string `dynamo:"imageUrl"`
	Title     string `dynamo:"title"`
	Memo      string `dynamo:"memo"`
	UpdatedAt int64  `dynamo:"updatedAt"`
}

func toModel(input entity.UTNAEntityFood, id int64, title, imageURL string) utnaFood {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	return utnaFood{
		ID:        id,
		URL:       input.URL,
		ImageURL:  imageURL,
		Title:     title,
		Memo:      input.Memo,
		UpdatedAt: time.Now().Unix(),
	}
}

func toEntity(input utnaFood) entity.UTNAEntityFood {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	return entity.UTNAEntityFood{
		ID:       input.ID,
		URL:      input.URL,
		ImageURL: input.ImageURL,
		Title:    input.Title,
		Memo:     input.Memo,
	}
}

package interactor

import (
	"context"
	"log"
	"time"

	"github.com/tsutarou10/line_project/service/pkg/entity"
	"github.com/tsutarou10/line_project/service/pkg/utils"
)

func (i *interactor) HandleVisit(ctx context.Context, url string) error {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	ogp := utils.FetchOGP(url)
	title := utils.FetchTitle(ogp)
	imageURL := utils.FetchImageURL(ogp)
	ogpTag := entity.OGPTag{
		URL:      url,
		Title:    title,
		ImageURL: imageURL,
	}

	if err := i.visited.Put(ctx, ogpTag); err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		return err
	}

	input := entity.UTNAEntityFood{
		URL:       url,
		Title:     title,
		ImageURL:  imageURL,
		Hidden:    true,
		UpdatedAt: time.Now().Unix(),
	}
	if err := i.utnaFood.Update(ctx, input); err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		return err
	}

	i.out.EmitVisit(ctx, url)
	return nil
}

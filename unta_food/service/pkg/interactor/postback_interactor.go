package interactor

import (
	"context"
	"log"

	"github.com/tsutarou10/line_project/service/pkg/entity"
	"github.com/tsutarou10/line_project/service/pkg/utils"
)

func (i *interactor) HandleComplete(ctx context.Context, url string) error {
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

	if err := i.dynamo.PutCompleted(ctx, ogpTag); err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		return err
	}
	i.out.EmitComplete(ctx, url)
	return nil
}

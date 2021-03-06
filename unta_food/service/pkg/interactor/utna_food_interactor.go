package interactor

import (
	"context"
	"log"
	"time"

	"github.com/tsutarou10/line_project/service/pkg/entity"
	"github.com/tsutarou10/line_project/service/pkg/utils"
)

func (i *interactor) HandleRegister(ctx context.Context, input entity.UTNAEntityFood) error {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	ogpTag := i.ogp.FetchOGPTag(input.URL)
	input.Title = ogpTag.Title
	input.ImageURL = ogpTag.ImageURL
	input.Hidden = false

	if err := i.utnaFood.Put(ctx, input); err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		return err
	}
	i.out.EmitRegister(ctx, input)
	return nil
}

func (i *interactor) HandleGetAll(ctx context.Context) error {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	res, err := i.utnaFood.GetAll(ctx)
	if err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		return err
	}

	i.out.EmitGetAll(ctx, res)
	return nil
}

func (i *interactor) HandleDelete(ctx context.Context, url string) error {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	res, err := i.utnaFood.Delete(ctx, url)
	if err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		return err
	}
	i.out.EmitDelete(ctx, *res)
	return nil
}

func (i *interactor) HandleUpdate(ctx context.Context, src entity.UTNAEntityFood) error {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	ogpTag := i.ogp.FetchOGPTag(src.URL)
	src.Title = ogpTag.Title
	src.ImageURL = ogpTag.ImageURL
	src.UpdatedAt = time.Now().Unix()
	src.Hidden = false

	if err := i.utnaFood.Update(ctx, src); err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		return err
	}
	i.out.EmitDelete(ctx, src)
	return nil
}

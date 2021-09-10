package dynamo

import (
	"context"
	"errors"
	"log"

	"github.com/tsutarou10/line_project/service/pkg/entity"
	"github.com/tsutarou10/line_project/service/pkg/utils"
)

func (d *Dynamo) Put(ctx context.Context, input entity.UTNAEntityFood) error {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	tmp, _ := d.getWithURL(ctx, input.URL)
	if tmp != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), "already registered")
		return errors.New("already registered")
	}

	ogp := utils.FetchOGP(input.URL)
	title := utils.FetchTitle(ogp)
	imageURL := utils.FetchImageURL(ogp)

	if err := d.utnaFood.Put(toModel(input, title, imageURL)).Run(); err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		return err
	}
	return nil
}

func (d *Dynamo) Update(ctx context.Context, input entity.UTNAEntityFood) error {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	ogp := utils.FetchOGP(input.URL)
	title := utils.FetchTitle(ogp)
	imageURL := utils.FetchImageURL(ogp)

	if err := d.utnaFood.Put(toModel(input, title, imageURL)).Run(); err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		return err
	}
	return nil
}

func (d *Dynamo) GetAll(ctx context.Context) ([]entity.UTNAEntityFood, error) {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	var resDynamo []utnaFood
	err := d.utnaFood.Scan().All(&resDynamo)
	if err != nil {
		return nil, err
	}

	var rsl []entity.UTNAEntityFood
	for _, r := range resDynamo {
		rsl = append(rsl, toEntity(r))
	}

	return rsl, nil
}

func (d *Dynamo) Delete(ctx context.Context, url string) (*entity.UTNAEntityFood, error) {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	var oldValue utnaFood
	err := d.utnaFood.Delete("url", url).OldValue(&oldValue)
	if err != nil {
		return nil, err
	}
	res := toEntity(oldValue)
	return &res, nil
}

func (d *Dynamo) getWithURL(ctx context.Context, url string) (*utnaFood, error) {
	var rsl utnaFood
	if err := d.utnaFood.Get("url", url).One(&rsl); err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		return nil, err
	}
	return &rsl, nil
}

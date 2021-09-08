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
	rs := d.getRegisterStatus(ctx)

	if err := d.utnaFood.Put(toModel(input, rs.Number+1)).Run(); err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		return err
	}
	return nil
}

func (d *Dynamo) Update(ctx context.Context, input entity.UTNAEntityFood) error {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())
	if err := d.utnaFood.Put(toModel(input, input.ID)).Run(); err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		return err
	}
	return nil
}

func (d *Dynamo) UpdateRegisterStatus(ctx context.Context, isAdd bool) error {
	rs := d.getRegisterStatus(ctx)

	if isAdd {
		rs.Number += 1
	} else {
		rs.Number -= 1
	}
	if err := d.putRegisterStatus(ctx, rs.Number); err != nil {
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

func (d *Dynamo) Delete(ctx context.Context, id int64) (*entity.UTNAEntityFood, error) {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	var oldValue utnaFood
	err := d.utnaFood.Delete("id", id).OldValue(&oldValue)
	if err != nil {
		return nil, err
	}
	res := toEntity(oldValue)
	return &res, nil
}

func (d *Dynamo) getWithURL(ctx context.Context, url string) (*utnaFood, error) {
	var rsl utnaFood
	if err := d.utnaFood.Get("url", url).Index("URLIndex").One(&rsl); err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		return nil, err
	}
	return &rsl, nil
}

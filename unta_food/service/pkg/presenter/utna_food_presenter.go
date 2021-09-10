package presenter

import (
	"context"
	"log"
	"sort"

	"github.com/tsutarou10/line_project/service/pkg/entity"
	"github.com/tsutarou10/line_project/service/pkg/utils"
)

type Presenter struct {
	registerCh chan interface{}
	getAllCh   chan interface{}
	deleteCh   chan interface{}
	completeCh chan interface{}
}

func NewPresenter() *Presenter {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	return &Presenter{
		make(chan interface{}),
		make(chan interface{}),
		make(chan interface{}),
		make(chan interface{}),
	}
}

func (p *Presenter) EmitRegister(ctx context.Context, output entity.UTNAEntityFood) {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	go func() {
		p.registerCh <- output
	}()
}

func (p *Presenter) WaitForRegisterCompleted(ctx context.Context) (interface{}, error) {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	return <-p.registerCh, nil
}

func (p *Presenter) EmitGetAll(ctx context.Context, output []entity.UTNAEntityFood) {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	go func() {
		var res []entity.UTNAEntityFood
		for _, o := range output {
			if !o.Hidden {
				res = append(res, o)
			}
		}
		sort.Slice(res, func(i, j int) bool {
			return res[i].UpdatedAt < res[j].UpdatedAt
		})
		p.getAllCh <- res
	}()
}

func (p *Presenter) WaitForGetAllCompleted(ctx context.Context) (interface{}, error) {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	return <-p.getAllCh, nil
}

func (p *Presenter) EmitDelete(ctx context.Context, output entity.UTNAEntityFood) {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	go func() {
		p.deleteCh <- output
	}()
}

func (p *Presenter) WaitForDeleteCompleted(ctx context.Context) (interface{}, error) {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	return <-p.deleteCh, nil
}

func (p *Presenter) EmitUpdate(ctx context.Context, output entity.UTNAEntityFood) {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	go func() {
		p.deleteCh <- output
	}()
}

func (p *Presenter) WaitForUpdateCompleted(ctx context.Context) (interface{}, error) {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	return <-p.deleteCh, nil
}

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
}

func NewPresenter() *Presenter {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	return &Presenter{
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
		sort.Slice(output, func(i, j int) bool {
			return output[i].UpdatedAt < output[j].UpdatedAt
		})
		p.getAllCh <- output
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

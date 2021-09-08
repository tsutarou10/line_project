package presenter

import (
	"context"
	"log"

	"github.com/tsutarou10/line_project/service/pkg/utils"
)

type Presenter struct {
	registerCh chan interface{}
}

func NewPresenter() *Presenter {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	return &Presenter{
		make(chan interface{}),
	}
}

func (p *Presenter) EmitRegister(ctx context.Context, output interface{}) {
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

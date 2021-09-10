package presenter

import (
	"context"
	"log"

	"github.com/tsutarou10/line_project/service/pkg/utils"
)

func (p *Presenter) EmitVisit(ctx context.Context, output string) {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	go func() {
		p.completeCh <- output
	}()
}

func (p *Presenter) WaitForVisitCompleted(ctx context.Context) (interface{}, error) {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	return <-p.completeCh, nil
}

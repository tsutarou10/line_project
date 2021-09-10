package presenter

import (
	"context"
	"log"
	"sort"

	"github.com/tsutarou10/line_project/service/pkg/entity"
	"github.com/tsutarou10/line_project/service/pkg/utils"
)

func (p *Presenter) EmitGetVisitedRestaurant(ctx context.Context, output []entity.UTNAEntityFood) {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	go func() {
		var res []entity.UTNAEntityFood
		for _, o := range output {
			res = append(res, o)
		}
		sort.Slice(res, func(i, j int) bool {
			return res[i].UpdatedAt < res[j].UpdatedAt
		})
		p.visitedCh <- res
	}()
}

func (p *Presenter) WaitForGetVisitedRestaurantCompleted(ctx context.Context) (interface{}, error) {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	return <-p.visitedCh, nil
}

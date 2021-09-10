package interactor

import (
	"context"
	"log"

	"github.com/tsutarou10/line_project/service/pkg/utils"
)

func (i *interactor) HandleGetVisitedRestaurant(ctx context.Context) error {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	res, err := i.visited.Scan(ctx)
	if err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		return err
	}

	i.out.EmitGetVisitedRestaurant(ctx, res)
	return nil
}

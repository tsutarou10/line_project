package interactor

import (
	"context"
	"log"

	"github.com/tsutarou10/line_project/service/pkg/utils"
)

func (i *interactor) HandleComplete(ctx context.Context, url string) error {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	if err := i.dynamo.PutCompleted(ctx, url); err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		return err
	}
	i.out.EmitComplete(ctx, url)
	return nil
}

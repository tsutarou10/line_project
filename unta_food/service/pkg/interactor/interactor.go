package interactor

import (
	"context"
	"log"

	"github.com/tsutarou10/line_project/service/pkg/entity"
	"github.com/tsutarou10/line_project/service/pkg/gateway"
	"github.com/tsutarou10/line_project/service/pkg/usecase"
	"github.com/tsutarou10/line_project/service/pkg/utils"
)

type interactor struct {
	out    usecase.OutputPort
	dynamo gateway.DynamoGateway
}

func NewInputPort(
	out usecase.OutputPort,
	dynamo gateway.DynamoGateway,
) usecase.InputPort {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	return &interactor{
		out,
		dynamo,
	}
}

func (i *interactor) HandleRegister(ctx context.Context, input entity.RegisterEntity) error {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	if err := i.dynamo.Put(ctx, input); err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		return err
	}
	i.out.EmitRegister(ctx, "success")
	return nil
}

func (i *interactor) HandleGetAll(ctx context.Context) error {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	res, err := i.dynamo.GetAll(ctx)
	if err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		return err
	}
	i.out.EmitGetAll(ctx, res)
	return nil
}

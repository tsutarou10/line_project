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
	if err := i.dynamo.UpdateRegisterStatus(ctx, true); err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		return err
	}
	i.out.EmitRegister(ctx, input)
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

func (i *interactor) HandleDelete(ctx context.Context, id int64) error {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	res, err := i.dynamo.Delete(ctx, id)
	if err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		return err
	}
	if err := i.dynamo.UpdateRegisterStatus(ctx, false); err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		return err
	}
	i.out.EmitDelete(ctx, *res)
	return nil
}

func (i *interactor) HandleUpdate(ctx context.Context, src entity.RegisterEntity) error {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	if err := i.dynamo.Update(ctx, src); err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		return err
	}
	i.out.EmitDelete(ctx, src)
	return nil
}

package usecase

import (
	"context"

	"github.com/tsutarou10/line_project/service/pkg/entity"
)

type UTNAFoodInputPort interface {
	HandleRegister(context.Context, entity.UTNAEntityFood) error
	HandleGetAll(context.Context) error
	HandleDelete(context.Context, string) error
	HandleUpdate(context.Context, entity.UTNAEntityFood) error
	HandleComplete(context.Context, string) error
}

type UTNAFoodOutputPort interface {
	EmitRegister(context.Context, entity.UTNAEntityFood)
	EmitGetAll(context.Context, []entity.UTNAEntityFood)
	EmitDelete(context.Context, entity.UTNAEntityFood)
	EmitUpdate(context.Context, entity.UTNAEntityFood)
	EmitComplete(context.Context, string)
}

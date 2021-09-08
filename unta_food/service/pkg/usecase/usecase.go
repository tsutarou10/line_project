package usecase

import (
	"context"

	"github.com/tsutarou10/line_project/service/pkg/entity"
)

type InputPort interface {
	HandleRegister(context.Context, entity.UTNAEntityFood) error
	HandleGetAll(context.Context) error
	HandleDelete(context.Context, int64) error
	HandleUpdate(context.Context, entity.UTNAEntityFood) error
}

type OutputPort interface {
	EmitRegister(context.Context, entity.UTNAEntityFood)
	EmitGetAll(context.Context, []entity.UTNAEntityFood)
	EmitDelete(context.Context, entity.UTNAEntityFood)
	EmitUpdate(context.Context, entity.UTNAEntityFood)
}

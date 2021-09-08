package usecase

import (
	"context"

	"github.com/tsutarou10/line_project/service/pkg/entity"
)

type InputPort interface {
	HandleRegister(context.Context, entity.RegisterEntity) error
	HandleGetAll(context.Context) error
	HandleDelete(context.Context, int64) error
	HandleUpdate(context.Context, entity.RegisterEntity) error
}

type OutputPort interface {
	EmitRegister(context.Context, entity.RegisterEntity)
	EmitGetAll(context.Context, []entity.RegisterEntity)
	EmitDelete(context.Context, entity.RegisterEntity)
	EmitUpdate(context.Context, entity.RegisterEntity)
}

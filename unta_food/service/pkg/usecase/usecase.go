package usecase

import (
	"context"

	"github.com/tsutarou10/line_project/service/pkg/entity"
)

type InputPort interface {
	HandleRegister(context.Context, entity.RegisterEntity) error
}

type OutputPort interface {
	EmitRegister(context.Context, interface{})
}

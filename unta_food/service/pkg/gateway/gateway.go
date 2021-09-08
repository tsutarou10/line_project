package gateway

import (
	"context"

	"github.com/tsutarou10/line_project/service/pkg/entity"
)

type DynamoGateway interface {
	Put(context.Context, entity.RegisterEntity) error
	GetAll(context.Context) ([]entity.RegisterEntity, error)
	Delete(context.Context, int64) (*entity.RegisterEntity, error)
	Update(context.Context, entity.RegisterEntity) error
	UpdateRegisterStatus(context.Context, bool) error
}

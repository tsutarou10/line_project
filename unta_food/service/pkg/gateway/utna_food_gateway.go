package gateway

import (
	"context"

	"github.com/tsutarou10/line_project/service/pkg/entity"
)

type UTNAFoodDynamoGateway interface {
	Put(context.Context, entity.UTNAEntityFood) error
	GetAll(context.Context) ([]entity.UTNAEntityFood, error)
	Delete(context.Context, string) (*entity.UTNAEntityFood, error)
	Update(context.Context, entity.UTNAEntityFood) error
}

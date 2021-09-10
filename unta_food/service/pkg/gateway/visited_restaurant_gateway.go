package gateway

import (
	"context"

	"github.com/tsutarou10/line_project/service/pkg/entity"
)

type VisitedRestaurantGateway interface {
	Put(context.Context, entity.OGPTag) error
	Scan(context.Context) ([]entity.UTNAEntityFood, error)
}

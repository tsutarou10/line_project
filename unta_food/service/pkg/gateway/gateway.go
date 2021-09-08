package gateway

import (
	"context"

	"github.com/tsutarou10/line_project/service/pkg/entity"
)

type DynamoGateway interface {
	Put(context.Context, entity.RegisterEntity) error
}

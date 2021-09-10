package interactor

import (
	"log"

	"github.com/tsutarou10/line_project/service/pkg/gateway"
	"github.com/tsutarou10/line_project/service/pkg/usecase"
	"github.com/tsutarou10/line_project/service/pkg/utils"
)

type interactor struct {
	out    usecase.UTNAFoodOutputPort
	dynamo gateway.DynamoGateway
}

func NewInputPort(
	out usecase.UTNAFoodOutputPort,
	dynamo gateway.DynamoGateway,
) usecase.UTNAFoodInputPort {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	return &interactor{
		out,
		dynamo,
	}
}

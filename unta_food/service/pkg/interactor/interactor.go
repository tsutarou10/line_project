package interactor

import (
	"log"

	"github.com/tsutarou10/line_project/service/pkg/gateway"
	"github.com/tsutarou10/line_project/service/pkg/usecase"
	"github.com/tsutarou10/line_project/service/pkg/utils"
)

type interactor struct {
	out      usecase.UTNAFoodOutputPort
	utnaFood gateway.UTNAFoodDynamoGateway
	visited  gateway.VisitedRestaurantGateway
	ogp      gateway.OpenGraphGateway
}

func NewInputPort(
	out usecase.UTNAFoodOutputPort,
	utnaFood gateway.UTNAFoodDynamoGateway,
	visited gateway.VisitedRestaurantGateway,
	ogp gateway.OpenGraphGateway,
) usecase.UTNAFoodInputPort {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	return &interactor{
		out,
		utnaFood,
		visited,
		ogp,
	}
}

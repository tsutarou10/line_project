package dynamo

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/tsutarou10/line_project/service/pkg/entity"
	"github.com/tsutarou10/line_project/service/pkg/utils"
)

type VisitedRestaurantDynamo struct {
	visited dynamo.Table
}

func NewVisitedRestaurantDynamo() *VisitedRestaurantDynamo {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	db := dynamo.New(
		session.New(),
		aws.NewConfig().
			WithRegion(os.Getenv("REGION")).
			WithEndpoint(os.Getenv("DYNAMODB_ENDPOINT")),
	)
	visited := db.Table(os.Getenv("VISITED_RESTAURANT_TABLE_NAME"))
	return &VisitedRestaurantDynamo{
		visited: visited,
	}
}

func (d *VisitedRestaurantDynamo) Put(ctx context.Context, ogpTag entity.OGPTag) error {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	input := entity.UTNAEntityFood{
		URL:       ogpTag.URL,
		Title:     ogpTag.Title,
		ImageURL:  ogpTag.ImageURL,
		UpdatedAt: time.Now().Unix(),
	}

	if err := d.visited.Put(toModelOfVisitedRestaurant(input)).Run(); err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		return err
	}
	return nil

}

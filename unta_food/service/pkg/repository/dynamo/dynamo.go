package dynamo

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/tsutarou10/line_project/service/pkg/utils"
)

type Dynamo struct {
	utnaFood dynamo.Table
}

func NewDynamo() *Dynamo {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	db := dynamo.New(
		session.New(),
		aws.NewConfig().
			WithRegion(os.Getenv("REGION")).
			WithEndpoint(os.Getenv("DYNAMODB_ENDPOINT")),
	)
	utnaFood := db.Table(os.Getenv("UTNA_FOOD_TABLE_NAME"))
	return &Dynamo{
		utnaFood: utnaFood,
	}
}

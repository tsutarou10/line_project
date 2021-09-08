package repository

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/tsutarou10/line_project/service/pkg/entity"
	"github.com/tsutarou10/line_project/service/pkg/utils"
)

type Dynamo struct {
	table dynamo.Table
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
	table := db.Table(os.Getenv("TABLE_NAME"))
	return &Dynamo{table: table}
}

func (d *Dynamo) Put(ctx context.Context, input entity.RegisterEntity) error {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	log.Print(input)
	log.Print(os.Getenv("DYNAMODB_ENDPOINT"))
	log.Print(os.Getenv("TABLE_NAME"))
	log.Print(os.Getenv("REGION"))
	if err := d.table.Put(toModel(input)).Run(); err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		return err
	}
	log.Print("end put")
	return nil
}

type testInput struct {
	URL string `dynamo:"url"`
}

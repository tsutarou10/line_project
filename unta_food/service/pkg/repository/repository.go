package repository

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/tsutarou10/line_project/service/pkg/entity"
	"github.com/tsutarou10/line_project/service/pkg/utils"
)

type Dynamo struct {
	utnaFood       dynamo.Table
	registerStatus dynamo.Table
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
	registerStatus := db.Table(os.Getenv("REGISTRATION_STATUS_TABLE_NAME"))
	return &Dynamo{
		utnaFood:       utnaFood,
		registerStatus: registerStatus,
	}
}

func (d *Dynamo) Put(ctx context.Context, input entity.RegisterEntity) error {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	tmp, _ := d.getWithURL(ctx, input.URL)
	if tmp != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), "already registered")
		return errors.New("already registered")
	}
	rs := d.getRegisterStatus(ctx)

	if err := d.utnaFood.Put(toModel(input, rs.Number+1)).Run(); err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		return err
	}
	if err := d.putRegisterStatus(ctx, rs.Number+1); err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		return err
	}
	return nil
}

func (d *Dynamo) GetAll(ctx context.Context) ([]entity.RegisterEntity, error) {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	var resDynamo []utnaFoodSchema
	err := d.utnaFood.Scan().All(&resDynamo)
	if err != nil {
		return nil, err
	}

	var rsl []entity.RegisterEntity
	for _, r := range resDynamo {
		rsl = append(rsl, toEntity(r))
	}

	return rsl, nil
}

func (d *Dynamo) getRegisterStatus(ctx context.Context) utnaFoodRegisterStatus {
	var res utnaFoodRegisterStatus
	err := d.registerStatus.Get("status", "registered").One(&res)
	if err != nil {
		return utnaFoodRegisterStatus{
			Status:    "registered",
			Number:    0,
			UpdatedAt: time.Now().Unix(),
		}
	}
	return res
}

func (d *Dynamo) putRegisterStatus(ctx context.Context, rn int64) error {
	input := utnaFoodRegisterStatus{
		Status:    "registered",
		Number:    rn,
		UpdatedAt: time.Now().Unix(),
	}
	if err := d.registerStatus.Put(input).Run(); err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		return err
	}
	return nil
}

func (d *Dynamo) getWithURL(ctx context.Context, url string) (*utnaFoodSchema, error) {
	var rsl utnaFoodSchema
	if err := d.utnaFood.Get("url", url).Index("URLIndex").One(&rsl); err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		return nil, err
	}
	return &rsl, nil
}

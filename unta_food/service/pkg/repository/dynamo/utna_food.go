package dynamo

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

type UTNAFoodDynamo struct {
	utnaFood dynamo.Table
}

func NewUTNAFoodDynamo() *UTNAFoodDynamo {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	db := dynamo.New(
		session.New(),
		aws.NewConfig().
			WithRegion(os.Getenv("REGION")).
			WithEndpoint(os.Getenv("DYNAMODB_ENDPOINT")),
	)
	utnaFood := db.Table(os.Getenv("UTNA_FOOD_TABLE_NAME"))
	return &UTNAFoodDynamo{
		utnaFood: utnaFood,
	}
}

func (d *UTNAFoodDynamo) Put(ctx context.Context, input entity.UTNAEntityFood) error {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	tmp, _ := d.getWithURL(ctx, input.URL)
	if tmp != nil {
		errMsg := "既に登録されています。情報更新する場合は update コマンドを用いてください。"
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), errMsg)
		return errors.New(errMsg)
	}

	input.UpdatedAt = time.Now().Unix()

	if err := d.utnaFood.Put(toModelOfUTNAFood(input)).Run(); err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		return err
	}
	return nil
}

func (d *UTNAFoodDynamo) Update(ctx context.Context, input entity.UTNAEntityFood) error {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	if err := d.utnaFood.Put(toModelOfUTNAFood(input)).Run(); err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		return err
	}
	return nil
}

func (d *UTNAFoodDynamo) GetAll(ctx context.Context) ([]entity.UTNAEntityFood, error) {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	var resDynamo []utnaFood
	err := d.utnaFood.Scan().All(&resDynamo)
	if err != nil {
		return nil, err
	}

	var rsl []entity.UTNAEntityFood
	for _, r := range resDynamo {
		rsl = append(rsl, toEntityOfUTNAFood(r))
	}

	return rsl, nil
}

func (d *UTNAFoodDynamo) Delete(ctx context.Context, url string) (*entity.UTNAEntityFood, error) {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	var oldValue utnaFood
	err := d.utnaFood.Delete("url", url).OldValue(&oldValue)
	if err != nil {
		return nil, err
	}
	res := toEntityOfUTNAFood(oldValue)
	return &res, nil
}

func (d *UTNAFoodDynamo) getWithURL(ctx context.Context, url string) (*utnaFood, error) {
	var rsl utnaFood
	if err := d.utnaFood.Get("url", url).One(&rsl); err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		return nil, err
	}
	return &rsl, nil
}

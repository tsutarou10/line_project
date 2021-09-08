package dynamo

import (
	"context"
	"log"
	"time"

	"github.com/tsutarou10/line_project/service/pkg/utils"
)

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

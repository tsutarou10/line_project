package controller

import (
	"context"
	"errors"
	"log"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/tsutarou10/line_project/service/pkg/entity"
	"github.com/tsutarou10/line_project/service/pkg/usecase"
	"github.com/tsutarou10/line_project/service/pkg/utils"
)

type Controller struct {
	in usecase.UTNAFoodInputPort
}

func NewController(in usecase.UTNAFoodInputPort) *Controller {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	return &Controller{in}
}

func (c *Controller) RegisterController(ctx context.Context, req events.APIGatewayProxyRequest) error {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	webhook, err := utils.ExtractWebhook(req)
	if err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		return err
	}

	wc := utils.ExtractWebhookContext(*webhook)
	if wc == nil || len(wc.ReceivedMessages) == 0 {
		msg := "invalid request"
		log.Printf("[ERROR]: %s. sm: %s", utils.GetFuncName(), msg)
		return errors.New(msg)
	}

	if !isURL(wc.ReceivedMessages[0]) {
		msg := "invalid url"
		log.Printf("[ERROR]: %s, %s is %s", utils.GetFuncName(), wc.ReceivedMessages[0], msg)
		return errors.New(msg)
	}
	input := entity.UTNAEntityFood{
		URL: wc.ReceivedMessages[0],
	}
	if len(wc.ReceivedMessages) > 1 {
		input.Memo = wc.ReceivedMessages[1]
	}
	return c.in.HandleRegister(ctx, input)
}

func (c *Controller) GetAllController(ctx context.Context, req events.APIGatewayProxyRequest) error {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	return c.in.HandleGetAll(ctx)
}

func (c *Controller) DeleteController(ctx context.Context, req events.APIGatewayProxyRequest) error {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	webhook, err := utils.ExtractWebhook(req)
	if err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		return err
	}

	wc := utils.ExtractWebhookContext(*webhook)
	if wc == nil || len(wc.ReceivedMessages) <= 1 {
		msg := "invalid requst"
		log.Printf("[ERROR]: %s. error: %s", utils.GetFuncName(), msg)
		return errors.New(msg)
	}

	id, err := strconv.ParseInt(wc.ReceivedMessages[1], 10, 64)
	if err != nil {
		log.Printf("[ERROR]: %s. error: %s", utils.GetFuncName(), "invalid id")
		return errors.New("invalid id")
	}

	return c.in.HandleDelete(ctx, id)

}

func (c *Controller) UpdateController(ctx context.Context, req events.APIGatewayProxyRequest) error {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	webhook, err := utils.ExtractWebhook(req)
	if err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		return err
	}

	wc := utils.ExtractWebhookContext(*webhook)
	// update id url [memo]
	if wc == nil || len(wc.ReceivedMessages) < 3 {
		msg := "invalid request"
		log.Printf("[ERROR]: %s. error: %s", utils.GetFuncName(), msg)
		return errors.New(msg)
	}

	if !isURL(wc.ReceivedMessages[2]) {
		msg := "invalid url"
		log.Printf("[ERROR]: %s, %s is %s", utils.GetFuncName(), wc.ReceivedMessages[0], msg)
		return errors.New(msg)
	}
	id, err := strconv.ParseInt(wc.ReceivedMessages[1], 10, 64)
	if err != nil {
		msg := "invalid id"
		log.Printf("[ERROR]: %s. error: %s", utils.GetFuncName(), msg)
		return errors.New(msg)
	}
	input := entity.UTNAEntityFood{
		ID:  id,
		URL: wc.ReceivedMessages[2],
	}
	if len(wc.ReceivedMessages) > 3 {
		input.Memo = wc.ReceivedMessages[3]
	}
	return c.in.HandleUpdate(ctx, input)
}

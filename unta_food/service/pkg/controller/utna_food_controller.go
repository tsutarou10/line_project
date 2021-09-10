package controller

import (
	"context"
	"errors"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/tsutarou10/line_project/service/pkg/entity"
	"github.com/tsutarou10/line_project/service/pkg/utils"
)

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

	if !utils.IsURL(wc.ReceivedMessages[0]) {
		msg := "invalid url"
		log.Printf("[ERROR]: %s, %s is %s", utils.GetFuncName(), wc.ReceivedMessages[0], msg)
		return errors.New(msg)
	}

	input := entity.UTNAEntityFood{
		URL: wc.ReceivedMessages[0],
	}

	if len(wc.ReceivedMessages) > 1 {
		input.Memo = createMemo(wc.ReceivedMessages[1:])
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

	return c.in.HandleDelete(ctx, wc.ReceivedMessages[1])
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
	// update url [memo]
	if wc == nil || len(wc.ReceivedMessages) < 2 {
		msg := "invalid request"
		log.Printf("[ERROR]: %s. error: %s", utils.GetFuncName(), msg)
		return errors.New(msg)
	}

	if !utils.IsURL(wc.ReceivedMessages[1]) {
		msg := "invalid url"
		log.Printf("[ERROR]: %s, %s is %s", utils.GetFuncName(), wc.ReceivedMessages[0], msg)
		return errors.New(msg)
	}
	input := entity.UTNAEntityFood{
		URL: wc.ReceivedMessages[1],
	}
	if len(wc.ReceivedMessages) > 2 {
		input.Memo = createMemo(wc.ReceivedMessages[3:])
	}
	return c.in.HandleUpdate(ctx, input)
}

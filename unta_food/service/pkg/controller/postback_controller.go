package controller

import (
	"context"
	"errors"
	"log"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/tsutarou10/line_project/service/pkg/utils"
)

func (c *Controller) DeleteControllerOfPostback(ctx context.Context, req events.APIGatewayProxyRequest) error {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	webhook, err := utils.ExtractWebhook(req)
	if err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		return err
	}

	wc := utils.ExtractWebhookContext(*webhook)
	if wc == nil {
		msg := "internal server error"
		log.Printf("[ERROR]: %s. error: %s", utils.GetFuncName(), msg)
		return errors.New(msg)
	}

	id, err := strconv.ParseInt(wc.ReceivedPostBackData["id"], 10, 64)
	if err != nil {
		errMsg := "internal server error"
		log.Printf("[ERROR]: %s. error: %s", utils.GetFuncName(), errMsg)
		return errors.New(errMsg)
	}

	return c.in.HandleDelete(ctx, id)

}

func (c *Controller) GetControllerOfPostback(ctx context.Context, req events.APIGatewayProxyRequest) error {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	webhook, err := utils.ExtractWebhook(req)
	if err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		return err
	}

	wc := utils.ExtractWebhookContext(*webhook)
	if wc == nil {
		msg := "internal server error"
		log.Printf("[ERROR]: %s. error: %s", utils.GetFuncName(), msg)
		return errors.New(msg)
	}

	return c.in.HandleGetAll(ctx)
}

package controller

import (
	"context"
	"errors"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/tsutarou10/line_project/service/pkg/entity"
	"github.com/tsutarou10/line_project/service/pkg/usecase"
	"github.com/tsutarou10/line_project/service/pkg/utils"
)

type Controller struct {
	in usecase.InputPort
}

func NewController(in usecase.InputPort) *Controller {
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
	if wc == nil || len(wc.ReceivedMessage) == 0 {
		log.Printf("[ERROR]: %s. sm: %s", utils.GetFuncName(), "error")
		return errors.New("error")
	}

	if !isURL(wc.ReceivedMessage) {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), "invalidate url")
		return errors.New("invalidate_url")
	}
	input := entity.RegisterEntity{
		URL: wc.ReceivedMessage,
	}
	return c.in.HandleRegister(ctx, input)
}

func isURL(s string) bool {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())
	return false

	//u, err := url.Parse(s)
	//return err == nil && u.Scheme != "" && u.Host != ""
}
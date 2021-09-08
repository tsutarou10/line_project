package controller

import (
	"context"
	"errors"
	"log"
	"net/url"
	"strconv"

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
	if wc == nil || len(wc.ReceivedMessages) == 0 {
		log.Printf("[ERROR]: %s. sm: %s", utils.GetFuncName(), "invalid request")
		return errors.New("invalid request")
	}

	if !isURL(wc.ReceivedMessages[0]) {
		log.Printf("[ERROR]: %s, %s is invalidate url", utils.GetFuncName(), wc.ReceivedMessages[0])
		return errors.New("invalidate url")
	}
	input := entity.RegisterEntity{
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
		log.Printf("[ERROR]: %s. sm: %s", utils.GetFuncName(), "invalid request")
		return errors.New("invalid request")
	}

	id, err := strconv.ParseInt(wc.ReceivedMessages[1], 10, 64)
	if err != nil {
		log.Printf("[ERROR]: %s. sm: %s", utils.GetFuncName(), "invalid id")
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
		log.Printf("[ERROR]: %s. sm: %s", utils.GetFuncName(), "invalid request")
		return errors.New("invalid request")
	}

	if !isURL(wc.ReceivedMessages[2]) {
		log.Printf("[ERROR]: %s, %s is invalidate url", utils.GetFuncName(), wc.ReceivedMessages[0])
		return errors.New("invalidate url")
	}
	id, err := strconv.ParseInt(wc.ReceivedMessages[1], 10, 64)
	if err != nil {
		log.Printf("[ERROR]: %s. sm: %s", utils.GetFuncName(), "invalid id")
		return errors.New("invalid id")
	}
	input := entity.RegisterEntity{
		ID:  id,
		URL: wc.ReceivedMessages[2],
	}
	if len(wc.ReceivedMessages) > 2 {
		input.Memo = wc.ReceivedMessages[3]
	}
	return c.in.HandleUpdate(ctx, input)
}

func isURL(s string) bool {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	u, err := url.Parse(s)
	return err == nil && u.Scheme != "" && u.Host != ""
}

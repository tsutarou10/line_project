package handler

import (
	"context"
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/tsutarou10/line_project/service/pkg/utils"
)

func createMethodPackageOfPostback(req events.APIGatewayProxyRequest) (*methodPackage, error) {
	wh, err := utils.ExtractWebhook(req)
	if err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		return nil, err
	}
	wc := utils.ExtractWebhookContext(*wh)
	var mp methodPackage
	mp.RequestType = "postback"
	switch strings.ToLower(wc.ReceivedPostBackData["action"]) {
	case "delete":
		mp.Foc = deleteHandlerOfPostback
		mp.Method = "delete"
	case "get":
		mp.Foc = getHandlerOfPostback
		mp.Method = "get"
	case "visit":
		mp.Foc = visitHandlerOfPostback
		mp.Method = "visit"
	case "update":
		mp.Foc = updateHandlerOfPostback
		mp.Method = "update"
	case "history":
		mp.Foc = getVisitedRestaurantOfPostback
		mp.Method = "history"
	}
	return &mp, nil
}

func deleteHandlerOfPostback(ctx context.Context, req events.APIGatewayProxyRequest) (interface{}, error) {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	c, p := setupAPIGatewayAdapter()
	log.Printf("%s, %s", utils.GetFuncName(), req.Body)
	if err := c.DeleteControllerOfPostback(ctx, req); err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		return nil, err
	}
	return p.WaitForDeleteCompleted(ctx)
}

func getHandlerOfPostback(ctx context.Context, req events.APIGatewayProxyRequest) (interface{}, error) {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	c, p := setupAPIGatewayAdapter()
	log.Printf("%s, %s", utils.GetFuncName(), req.Body)
	if err := c.GetControllerOfPostback(ctx, req); err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		return nil, err
	}
	return p.WaitForGetAllCompleted(ctx)
}

func visitHandlerOfPostback(ctx context.Context, req events.APIGatewayProxyRequest) (interface{}, error) {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	c, p := setupAPIGatewayAdapter()
	log.Printf("%s, %s", utils.GetFuncName(), req.Body)
	if err := c.VisitControllerOfPostback(ctx, req); err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		return nil, err
	}
	return p.WaitForVisitCompleted(ctx)
}

func updateHandlerOfPostback(ctx context.Context, req events.APIGatewayProxyRequest) (interface{}, error) {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	c, p := setupAPIGatewayAdapter()
	log.Printf("%s, %s", utils.GetFuncName(), req.Body)
	if err := c.UpdateControllerOfPostback(ctx, req); err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		return nil, err
	}
	return p.WaitForUpdateCompleted(ctx)
}

func getVisitedRestaurantOfPostback(ctx context.Context, req events.APIGatewayProxyRequest) (interface{}, error) {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	c, p := setupAPIGatewayAdapter()
	log.Printf("%s, %s", utils.GetFuncName(), req.Body)
	if err := c.GetVisitedRestaurantController(ctx, req); err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		return nil, err
	}
	return p.WaitForGetVisitedRestaurantCompleted(ctx)
}

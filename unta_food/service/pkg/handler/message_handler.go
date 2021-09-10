package handler

import (
	"context"
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/tsutarou10/line_project/service/pkg/utils"
)

var (
	README_URL = "https://github.com/tsutarou10/line_project/blob/main/unta_food/README.md"
)

func createMethodPackageOfMessage(req events.APIGatewayProxyRequest) (*methodPackage, error) {
	wh, err := utils.ExtractWebhook(req)
	if err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		return nil, err
	}
	wc := utils.ExtractWebhookContext(*wh)
	var mp methodPackage
	mp.RequestType = "message"
	switch strings.ToLower(wc.ReceivedMessages[0]) {
	case "get":
		mp.Foc = getAllHandlerOfMessage
		mp.Method = "get"
	case "delete":
		mp.Foc = deleteHandlerOfMessage
		mp.Method = "delete"
	case "update":
		mp.Foc = updateHandlerOfMessage
		mp.Method = "update"
	case "help", "ヘルプ":
		mp.Foc = printHelpOfMessage
		mp.Method = "help"
	default:
		mp.Foc = registerHandlerOfMessage
		mp.Method = "register"
	}
	return &mp, nil
}

func updateHandlerOfMessage(ctx context.Context, req events.APIGatewayProxyRequest) (interface{}, error) {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	c, p := setupAPIGatewayAdapter()
	log.Printf("%s, %s", utils.GetFuncName(), req.Body)
	if err := c.UpdateController(ctx, req); err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		return nil, err
	}
	return p.WaitForUpdateCompleted(ctx)
}

func registerHandlerOfMessage(ctx context.Context, req events.APIGatewayProxyRequest) (interface{}, error) {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	c, p := setupAPIGatewayAdapter()
	log.Printf("%s, %s", utils.GetFuncName(), req.Body)
	if err := c.RegisterController(ctx, req); err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		return nil, err
	}
	return p.WaitForRegisterCompleted(ctx)
}

func printHelpOfMessage(ctx context.Context, req events.APIGatewayProxyRequest) (interface{}, error) {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	return nil, nil
}

func getAllHandlerOfMessage(ctx context.Context, req events.APIGatewayProxyRequest) (interface{}, error) {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	c, p := setupAPIGatewayAdapter()
	log.Printf("%s, %s", utils.GetFuncName(), req.Body)
	if err := c.GetAllController(ctx, req); err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		return nil, err
	}
	return p.WaitForGetAllCompleted(ctx)
}

func deleteHandlerOfMessage(ctx context.Context, req events.APIGatewayProxyRequest) (interface{}, error) {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	c, p := setupAPIGatewayAdapter()
	log.Printf("%s, %s", utils.GetFuncName(), req.Body)
	if err := c.DeleteController(ctx, req); err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		return nil, err
	}
	return p.WaitForDeleteCompleted(ctx)
}

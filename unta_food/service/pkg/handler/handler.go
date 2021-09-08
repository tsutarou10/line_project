package handler

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/tsutarou10/line_project/service/pkg/controller"
	"github.com/tsutarou10/line_project/service/pkg/entity"
	"github.com/tsutarou10/line_project/service/pkg/interactor"
	"github.com/tsutarou10/line_project/service/pkg/presenter"
	"github.com/tsutarou10/line_project/service/pkg/repository"
	"github.com/tsutarou10/line_project/service/pkg/utils"
)

type methodPackage struct {
	Foc func(context.Context, events.APIGatewayProxyRequest) (interface{}, error)
}

func NewHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	if !validateSignature(os.Getenv("LINE_BOT_CHANNEL_SECRET"), request.Headers["x-line-signature"], []byte(request.Body)) {
		return raiseHandlerError(http.StatusUnauthorized, errors.New("invalidate signature"), request)
	}

	mp, err := createMethodPackage(request)
	if err != nil {
		return raiseHandlerError(500, err, request)
	}
	out, err := mp.Foc(ctx, request)
	if err != nil {
		return raiseHandlerError(500, err, request)
	}

	msg := convertReplyMessage(out)
	utils.ReplyMessageUsingAPIGWRequest(request, msg)
	return newAPIGatewayProxyReseponse(201, nil, request), nil
}

func newAPIGatewayProxyReseponse(statusCode int, err error, request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	if statusCode != 200 {
		utils.ReplyMessageUsingAPIGWRequest(request, err.Error())
	} else {
		utils.ReplyMessageUsingAPIGWRequest(request, "success")
	}

	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       fmt.Sprintf(`{"message:""$s"}`+"\n", http.StatusText(statusCode)),
	}
}

func validateSignature(channelSecret string, signature string, body []byte) bool {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	decoded, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false
	}
	hash := hmac.New(sha256.New, []byte(channelSecret))
	_, err = hash.Write(body)
	if err != nil {
		return false
	}
	return hmac.Equal(decoded, hash.Sum(nil))
}

func registerHandler(ctx context.Context, req events.APIGatewayProxyRequest) (interface{}, error) {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	c, p := setupAPIGatewayAdapter()
	log.Printf("%s, %s", utils.GetFuncName(), req.Body)
	if err := c.RegisterController(ctx, req); err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		utils.ReplyMessageUsingAPIGWRequest(req, err.Error())
		return nil, err
	}
	return p.WaitForRegisterCompleted(ctx)
}

func printHelp(ctx context.Context, req events.APIGatewayProxyRequest) (interface{}, error) {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	wh, err := utils.ExtractWebhook(req)
	if err != nil {
		raiseHandlerError(500, err, req)
	}
	wc := utils.ExtractWebhookContext(*wh)
	msg := `get: 登録された飲食店の URL とメモを取得できます。
URL メモ: 飲食店の URL とそのメモを登録できます。（メモは任意)
`
	return nil, utils.ReplyMessage(*wc, msg)
}

func getAllHandler(ctx context.Context, req events.APIGatewayProxyRequest) (interface{}, error) {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	c, p := setupAPIGatewayAdapter()
	log.Printf("%s, %s", utils.GetFuncName(), req.Body)
	if err := c.GetAllController(ctx, req); err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		utils.ReplyMessageUsingAPIGWRequest(req, err.Error())
		return nil, err
	}
	return p.WaitForGetAllCompleted(ctx)
}

func setupAPIGatewayAdapter() (*controller.Controller, *presenter.Presenter) {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	p := presenter.NewPresenter()
	dynamo := repository.NewDynamo()
	c := controller.NewController(
		interactor.NewInputPort(
			p,
			dynamo,
		))
	return c, p
}

func createMethodPackage(req events.APIGatewayProxyRequest) (*methodPackage, error) {
	wh, err := utils.ExtractWebhook(req)
	if err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		return nil, err
	}
	wc := utils.ExtractWebhookContext(*wh)
	var mp methodPackage
	switch wc.ReceivedMessages[0] {
	case "get":
		mp.Foc = getAllHandler
	case "help", "ヘルプ":
		mp.Foc = printHelp
	default:
		mp.Foc = registerHandler
	}
	return &mp, nil
}

func raiseHandlerError(statusCode int, err error, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
	utils.ReplyMessageUsingAPIGWRequest(req, err.Error())
	return newAPIGatewayProxyReseponse(statusCode, err, req), err
}

func convertReplyMessage(src interface{}) string {
	res := ""
	switch s := src.(type) {
	case []entity.RegisterEntity:
		for _, element := range s {
			if element.Memo != "" {
				res += fmt.Sprintf("・ %d: %s | %s\n", element.ID, element.URL, element.Memo)
			} else {
				res += fmt.Sprintf("・ %s\n", element.URL)
			}
		}
	case entity.RegisterEntity:
		res = s.URL
	default:
		res = "success"
	}
	return res
}

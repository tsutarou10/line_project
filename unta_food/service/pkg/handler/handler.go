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
	"github.com/tsutarou10/line_project/service/pkg/interactor"
	"github.com/tsutarou10/line_project/service/pkg/presenter"
	"github.com/tsutarou10/line_project/service/pkg/repository"
	"github.com/tsutarou10/line_project/service/pkg/utils"
)

func NewHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	if !validateSignature(os.Getenv("LINE_BOT_CHANNEL_SECRET"), request.Headers["x-line-signature"], []byte(request.Body)) {
		log.Printf("[ERROR]: %s, invalidate signature", utils.GetFuncName())
		utils.ReplyMessageUsingAPIGWRequest(request, "invalidate signature")
		return newAPIGatewayProxyReseponse(http.StatusBadRequest, errors.New("invalidate signature"), request), nil
	}

	_, err := registerHandler(ctx, request)
	if err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		utils.ReplyMessageUsingAPIGWRequest(request, err.Error())
		return newAPIGatewayProxyReseponse(500, err, request), err
	}
	utils.ReplyMessageUsingAPIGWRequest(request, "success")
	return newAPIGatewayProxyReseponse(200, nil, request), nil
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

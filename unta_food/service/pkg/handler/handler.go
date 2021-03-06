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
	"github.com/tsutarou10/line_project/service/pkg/repository/dynamo"
	"github.com/tsutarou10/line_project/service/pkg/repository/opengraph"
	"github.com/tsutarou10/line_project/service/pkg/utils"
)

type methodPackage struct {
	Foc         func(context.Context, events.APIGatewayProxyRequest) (interface{}, error)
	RequestType string
	Method      string
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

	if err = replyMessage(request, *mp, out); err != nil {
		return newAPIGatewayProxyReseponse(500, err, request), nil
	}
	return newAPIGatewayProxyReseponse(200, nil, request), nil
}

func createMethodPackage(req events.APIGatewayProxyRequest) (*methodPackage, error) {
	wh, err := utils.ExtractWebhook(req)
	if err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		return nil, err
	}
	wc := utils.ExtractWebhookContext(*wh)
	if len(wc.ReceivedMessages) != 0 {
		return createMethodPackageOfMessage(req)
	} else if len(wc.ReceivedPostBackData) != 0 {
		return createMethodPackageOfPostback(req)
	}

	errMsg := "invalid method"
	log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), errMsg)
	return nil, errors.New(errMsg)
}

func newAPIGatewayProxyReseponse(statusCode int, err error, request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

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

func setupAPIGatewayAdapter() (*controller.Controller, *presenter.Presenter) {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	p := presenter.NewPresenter()
	utnaFood := dynamo.NewUTNAFoodDynamo()
	visited := dynamo.NewVisitedRestaurantDynamo()
	ogp := opengraph.NewOpenGraph()
	c := controller.NewController(
		interactor.NewInputPort(
			p,
			utnaFood,
			visited,
			ogp,
		))
	return c, p
}

func raiseHandlerError(statusCode int, err error, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
	utils.ReplyMessageWithQuickResponse(req, err.Error(), README_URL, "LINE Bot???????????????")
	return newAPIGatewayProxyReseponse(statusCode, err, req), err
}

func replyMessage(req events.APIGatewayProxyRequest, mp methodPackage, src interface{}) error {
	switch mp.RequestType {
	case "message":
		return replyMessageOfMessage(req, mp, src)
	case "postback":
		return replyMessageOfPostback(req, mp, src)
	}
	errMsg := "invalid request type"
	if err := utils.ReplyMessageWithQuickResponse(req, errMsg, README_URL, "LINE Bot???????????????"); err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		return err
	}
	return errors.New(errMsg)
}

package utils

import (
	"encoding/json"
	"errors"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/line/line-bot-sdk-go/linebot"
)

type Webhook struct {
	Destination string           `json:"destination"`
	Events      []*linebot.Event `json:"events"`
}

type WebhookContext struct {
	ReceivedMessage string
	ReplyToken      string
}

func ExtractWebhookContext(wh Webhook) *WebhookContext {
	log.Printf("[START] :%s", GetFuncName())
	defer log.Printf("[END] :%s", GetFuncName())

	for _, event := range wh.Events {
		switch event.Type {
		case linebot.EventTypeMessage:
			switch m := event.Message.(type) {
			case *linebot.TextMessage:
				return &WebhookContext{
					ReceivedMessage: m.Text,
					ReplyToken:      event.ReplyToken,
				}
			}
		}
	}
	return nil
}

func ReplyMessage(wc WebhookContext, msg string) error {
	log.Printf("[START] :%s", GetFuncName())
	defer log.Printf("[END] :%s", GetFuncName())

	bot, err := linebot.New(
		os.Getenv("LINE_BOT_CHANNEL_SECRET"),
		os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Printf("[ERROR]: %s, %s", GetFuncName(), err.Error())
		return err
	}
	log.Printf("ReplyToken: %s", wc.ReplyToken)
	log.Printf("Message: %s", msg)
	bot.ReplyMessage(wc.ReplyToken, linebot.NewTextMessage(msg)).Do()
	return nil
}

func ExtractWebhook(req events.APIGatewayProxyRequest) (*Webhook, error) {
	log.Printf("[START] :%s", GetFuncName())
	defer log.Printf("[END] :%s", GetFuncName())

	var wh Webhook
	log.Print(req.Body)
	if err := json.Unmarshal([]byte(req.Body), &wh); err != nil {
		log.Printf("[ERROR]: %s, %s", GetFuncName(), err.Error())
		return nil, err
	}
	return &wh, nil
}

func ReplyMessageUsingAPIGWRequest(req events.APIGatewayProxyRequest, msg string) error {
	log.Printf("[START] :%s", GetFuncName())
	defer log.Printf("[END] :%s", GetFuncName())

	wh, err := ExtractWebhook(req)
	if err != nil {
		log.Printf("[ERROR]: %s, %s", GetFuncName(), err.Error())
		return err
	}
	wc := ExtractWebhookContext(*wh)
	if wc == nil {
		log.Printf("[ERROR]: %s, %s", GetFuncName(), "error")
		return errors.New("error")
	}
	if err = ReplyMessage(*wc, msg); err != nil {
		log.Printf("[ERROR]: %s, %s", GetFuncName(), "error")
		return errors.New("error")
	}
	return nil
}

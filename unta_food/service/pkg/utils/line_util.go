package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/tsutarou10/line_project/service/pkg/entity"
)

type Webhook struct {
	Destination string           `json:"destination"`
	Events      []*linebot.Event `json:"events"`
}

type WebhookContext struct {
	ReceivedMessages []string
	ReplyToken       string
}

func ExtractWebhookContext(wh Webhook) *WebhookContext {
	log.Printf("[START] :%s", GetFuncName())
	defer log.Printf("[END] :%s", GetFuncName())

	for _, event := range wh.Events {
		switch event.Type {
		case linebot.EventTypeMessage:
			switch m := event.Message.(type) {
			case *linebot.TextMessage:
				texts := SplitMultiSep(m.Text, []string{" ", "\n", "　"})
				return &WebhookContext{
					ReceivedMessages: texts,
					ReplyToken:       event.ReplyToken,
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

func ReplyCurousel(req events.APIGatewayProxyRequest, wc WebhookContext, src []entity.UTNAEntityFood) {
	bot, _ := linebot.New(
		os.Getenv("LINE_BOT_CHANNEL_SECRET"),
		os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
	)

	var ccList []*linebot.CarouselColumn
	for _, s := range src {
		cc := CreateCarouselColumn(s.Memo, s.URL, s.ID)
		ccList = append(ccList, cc)
	}
	ct := CreateCarouselTemplate(ccList)
	resp := linebot.NewTemplateMessage(
		"this is a carousel template with imageAspectRatio",
		ct,
	)
	b, err := bot.ReplyMessage(wc.ReplyToken, resp).Do()
	log.Print(b)
	log.Print(err)
}

func CreateCarouselTemplate(columns []*linebot.CarouselColumn) *linebot.CarouselTemplate {
	return linebot.NewCarouselTemplate(columns...).WithImageOptions("rectangle", "cover")
}

func CreateCarouselColumn(memo, url string, id int64) *linebot.CarouselColumn {

	description := fmt.Sprintf("%d", id)
	if memo != "" {
		description += fmt.Sprintf(" %s", memo)
	} else {
		description += fmt.Sprintf(" %s", url)
	}
	return linebot.NewCarouselColumn(
		"",
		"",
		description,
		linebot.NewURIAction("View detail", url),
		linebot.NewPostbackAction("Delete", fmt.Sprintf("action=delete&id=%d", id), "", ""),
	)
}

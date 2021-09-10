package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/tsutarou10/line_project/service/pkg/entity"
)

var (
	SAMPLE_IMAGE_URL = "https://utna-food-artifact-bucket.s3.ap-northeast-1.amazonaws.com/assets/images/rice.jpeg"
)

type Webhook struct {
	Destination string           `json:"destination"`
	Events      []*linebot.Event `json:"events"`
}

type WebhookContext struct {
	ReceivedMessages     []string
	ReceivedPostBackData map[string]string
	ReplyToken           string
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
		case linebot.EventTypePostback:
			// e.g. "action=buy&itemid=111"
			datas := strings.Split(event.Postback.Data, "&")
			rpbd := map[string]string{}
			for _, data := range datas {
				tmp := strings.Split(data, "=")
				rpbd[tmp[0]] = tmp[1]
			}
			return &WebhookContext{
				ReceivedPostBackData: rpbd,
				ReplyToken:           event.ReplyToken,
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
		cc := CreateCarouselColumn(s.Title, s.Memo, s.URL, s.ImageURL)
		ccList = append(ccList, cc)
	}
	ct := CreateCarouselTemplate(ccList)
	resp := linebot.NewTemplateMessage(
		"this is a carousel template with imageAspectRatio",
		ct,
	)
	out, err := bot.ReplyMessage(wc.ReplyToken, resp).Do()
	log.Printf("[RESP]: %v", out)
	log.Printf("[ERROR]: %v", err)
}

func CreateCarouselTemplate(columns []*linebot.CarouselColumn) *linebot.CarouselTemplate {
	return linebot.NewCarouselTemplate(columns...).WithImageOptions("rectangle", "cover")
}

func CreateCarouselColumn(title, memo, url, imageURL string) *linebot.CarouselColumn {
	if len(title) == 0 {
		title = "Empty title"
	}
	if len(title) > 40 {
		title = title[:40]
	}
	if len(imageURL) == 0 {
		imageURL = SAMPLE_IMAGE_URL
	}

	description := memo
	if description == "" {
		description += "Empty memo"
	}
	return linebot.NewCarouselColumn(
		imageURL,
		title,
		description,
		linebot.NewURIAction("View detail", url),
		linebot.NewPostbackAction("Delete", fmt.Sprintf("action=delete&url=%s", url), "", ""),
	)
}

func CreateQuickResponse(msg, url, buttonMsg string) linebot.SendingMessage {
	resp := linebot.NewTextMessage(
		msg,
	).WithQuickReplies(
		linebot.NewQuickReplyItems(
			linebot.NewQuickReplyButton(
				"",
				linebot.NewPostbackAction("登録情報を取得", "action=get", "", ""),
			),
			linebot.NewQuickReplyButton(
				"",
				linebot.NewURIAction(buttonMsg, url),
			),
		),
	)
	return resp
}

func ReplyMessageWithQuickResponse(req events.APIGatewayProxyRequest, msg, url, buttonMsg string) error {
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
	qr := CreateQuickResponse(msg, url, buttonMsg)
	bot, _ := linebot.New(
		os.Getenv("LINE_BOT_CHANNEL_SECRET"),
		os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
	)
	_, err = bot.ReplyMessage(wc.ReplyToken, qr).Do()
	return err
}

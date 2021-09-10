package handler

import (
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/tsutarou10/line_project/service/pkg/entity"
	"github.com/tsutarou10/line_project/service/pkg/utils"
)

const HELP_MESSAGE = `・get: 登録されてから行っていない飲食店の情報を取得できます。
  一度行ったことのある飲食店の情報は update コマンドで更新しないと表示されないので注意してください。
 
・飲食店のURL メモ: 飲食店の情報とそのメモを登録できます。（メモは任意)
  既に登録されている飲食店の情報を更新したい場合は update コマンドを使ってください。
 
・update URL メモ: 登録されている飲食店の情報 (URL先) を更新できます。
  登録されていない場合は新規登録します。(メモは任意)
  既に行ったことある飲食店に再度行きたくなった場合も update コマンドを用いて再度情報更新してください。
		`

func replyMessageOfMessage(req events.APIGatewayProxyRequest, mp methodPackage, src interface{}) error {
	msg := ""
	switch mp.Method {
	case "get":
		s := src.([]entity.UTNAEntityFood)
		wh, err := utils.ExtractWebhook(req)
		if err != nil {
			raiseHandlerError(500, err, req)
		}
		wc := utils.ExtractWebhookContext(*wh)
		utils.ReplyCurousel(req, *wc, s)
		return nil
	case "register":
		s := src.(entity.UTNAEntityFood)
		if s.Memo != "" {
			msg = fmt.Sprintf("registered %s | %s", s.URL, s.Memo)
		} else {
			msg = fmt.Sprintf("registered %s", s.URL)
		}
	case "update":
		s := src.(entity.UTNAEntityFood)
		if s.Memo != "" {
			msg = fmt.Sprintf("updated %s | %s", s.URL, s.Memo)
		} else {
			msg = fmt.Sprintf("updated %s", s.URL)
		}
	case "delete":
		s := src.(entity.UTNAEntityFood)
		if s.Memo != "" {
			msg = fmt.Sprintf("deleted %s | %s", s.URL, s.Memo)
		} else {
			msg = fmt.Sprintf("deleted %s", s.URL)
		}
	case "help":
		msg = HELP_MESSAGE
	default:
		msg = "Empty message"
	}
	if err := utils.ReplyMessageWithQuickResponse(req, msg, README_URL, "LINE Botの使用方法"); err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		return err
	}
	return nil
}

func replyMessageOfPostback(req events.APIGatewayProxyRequest, mp methodPackage, src interface{}) error {
	msg := ""
	switch mp.Method {
	case "get":
		s := src.([]entity.UTNAEntityFood)
		wh, err := utils.ExtractWebhook(req)
		if err != nil {
			raiseHandlerError(500, err, req)
		}
		wc := utils.ExtractWebhookContext(*wh)
		utils.ReplyCurousel(req, *wc, s)
		return nil
	case "update":
		s := src.(entity.UTNAEntityFood)
		if s.Memo != "" {
			msg = fmt.Sprintf("updated %s | %s", s.URL, s.Memo)
		} else {
			msg = fmt.Sprintf("updated %s", s.URL)
		}
	case "delete":
		s := src.(entity.UTNAEntityFood)
		if s.Memo != "" {
			msg = fmt.Sprintf("deleted %s | %s", s.URL, s.Memo)
		} else {
			msg = fmt.Sprintf("deleted %s", s.URL)
		}
	case "visit":
		msg = "good works!"
	case "help":
		msg = HELP_MESSAGE
	default:
		msg = "Empty message"
	}
	if err := utils.ReplyMessageWithQuickResponse(req, msg, README_URL, "LINE Botの使用方法"); err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		return err
	}
	return nil
}

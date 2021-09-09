package handler

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/tsutarou10/line_project/service/pkg/entity"
	"github.com/tsutarou10/line_project/service/pkg/utils"
)

func updateHandlerOfMessage(ctx context.Context, req events.APIGatewayProxyRequest) (interface{}, error) {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())

	c, p := setupAPIGatewayAdapter()
	log.Printf("%s, %s", utils.GetFuncName(), req.Body)
	if err := c.UpdateController(ctx, req); err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		utils.ReplyMessageUsingAPIGWRequest(req, err.Error())
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
		utils.ReplyMessageUsingAPIGWRequest(req, err.Error())
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
		utils.ReplyMessageUsingAPIGWRequest(req, err.Error())
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
		utils.ReplyMessageUsingAPIGWRequest(req, err.Error())
		return nil, err
	}
	return p.WaitForDeleteCompleted(ctx)
}

func createMethodPackageOfMessage(req events.APIGatewayProxyRequest) (*methodPackage, error) {
	wh, err := utils.ExtractWebhook(req)
	if err != nil {
		log.Printf("[ERROR]: %s, %s", utils.GetFuncName(), err.Error())
		return nil, err
	}
	wc := utils.ExtractWebhookContext(*wh)
	var mp methodPackage
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

func replyMessage(req events.APIGatewayProxyRequest, mp methodPackage, src interface{}) error {
	msg := ""
	switch s := src.(type) {
	case []entity.UTNAEntityFood:
		wh, err := utils.ExtractWebhook(req)
		if err != nil {
			raiseHandlerError(500, err, req)
		}
		wc := utils.ExtractWebhookContext(*wh)
		utils.ReplyCurousel(req, *wc, s)
	case entity.UTNAEntityFood:
		switch mp.Method {
		case "delete":
			if s.Memo != "" {
				msg = fmt.Sprintf("deleted %s | %s", s.URL, s.Memo)
			} else {
				msg = fmt.Sprintf("deleted %s", s.URL)
			}
		case "update":
			if s.Memo != "" {
				msg = fmt.Sprintf("updated %s | %s", s.URL, s.Memo)
			} else {
				msg = fmt.Sprintf("updated %s", s.URL)
			}
		case "register":
			if s.Memo != "" {
				msg = fmt.Sprintf("registered %s | %s", s.URL, s.Memo)
			} else {
				msg = fmt.Sprintf("registered %s", s.URL)
			}
		default:
			msg = s.URL
		}
		utils.ReplyMessageUsingAPIGWRequest(req, msg)
	default:
		switch mp.Method {
		case "help":
			msg = `・get: 登録された飲食店の URL とメモを取得できます。
	表示例 => id: URL | メモ

・URL メモ: 飲食店の URL とそのメモを登録できます。（メモは任意)

・update id url メモ: id で登録されている飲食店の情報を更新できます。id は get コマンドで確認できます。(メモは任意)

・delete id: 該当する id の飲食店を削除します。id は get コマンドで確認できます。

・詳細はこちら -> https://github.com/tsutarou10/line_project/blob/main/unta_food/README.md
		`
		default:
			msg = "success"
		}
		utils.ReplyMessageUsingAPIGWRequest(req, msg)
	}
	return nil
}

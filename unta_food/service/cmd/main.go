package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/tsutarou10/line_project/service/pkg/handler"
	"github.com/tsutarou10/line_project/service/pkg/utils"
)

func start(ctx context.Context, request events.APIGatewayProxyRequest) (interface{}, error) {
	return handler.NewHandler(ctx, request)
}

func main() {
	log.Printf("[START] :%s", utils.GetFuncName())
	defer log.Printf("[END] :%s", utils.GetFuncName())
	lambda.Start(start)
}

.PHONY: deps clean build

all: deps clean build

deps:
	go get -u github.com/aws/aws-lambda-go/events
	go get -u github.com/aws/aws-lambda-go/lambda

clean:
	rm -rf ./cmd/main

build:
	GOOS=linux GOARCH=amd64 go build -o ./cmd/main ./cmd

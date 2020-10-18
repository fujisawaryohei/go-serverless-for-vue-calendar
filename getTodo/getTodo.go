package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Request struct {
	id        string
	timestamp string
}

type Response struct {
	data string
}

func getTodo(request Request) (*dynamodb.GetItemOutput, error) {
	mySession := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1")},
	))
	svc := dynamodb.New(mySession)

	getTodoInput := &dynamodb.GetItemInput{
		TableName: aws.String("my-vue-calendar-db"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(request.id),
			},
			"timestamp": {
				S: aws.String(request.timestamp),
			},
		},
	}

	getItem, getErr := svc.GetItem(getTodoInput)
	if getErr != nil {
		panic(getErr)
	}
	return getItem, nil
}

func main() {
	lambda.Start(getTodo)
}

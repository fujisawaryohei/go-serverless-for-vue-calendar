package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Request events.APIGatewayProxyRequest
type Response events.APIGatewayProxyResponse

type Item struct {
	Timestamp string `json: timestamp`
	Content   string `json: content`
}

func Hanlder(ctx context.Context, request Request) (Response, error) {
	mySession := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1")},
	))
	svc := dynamodb.New(mySession)

	item := Item{}
	if err := json.Unmarshal([]uint8(request.Body), &item); err != nil {
		panic(err)
	}

	putItemInput := &dynamodb.PutItemInput{
		TableName: aws.String("my-vue-calendar-db"),
		Item: map[string]*dynamodb.AttributeValue{
			"timestamp": {
				S: aws.String(item.Timestamp),
			},
			"content": {
				S: aws.String(item.Content),
			},
		},
	}

	_, err := svc.PutItem(putItemInput)
	if err != nil {
		panic(err)
	}

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            "success",
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
		},
	}
	return resp, nil
}

func main() {
	lambda.Start(Hanlder)
}

package main

import (
	"fmt"
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

//https://godoc.org/github.com/aws/aws-lambda-go/events
func Handler(ctx context.Context, request Request) (Response, error) {
	mySession := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1")},
	))
	svc := dynamodb.New(mySession)
	//TODO: ↓のリファレンスを参照にQueryを実装する
	// https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/#QueryInput
	queryInput := &dynamodb.QueryInput{
		TableName: aws.String("my-vue-calendar-db"),
		KeyConditionExpression: aws.String("#timestamp = :timestamp"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			"timestamp": {
				S: aws.String(request.QueryStringParameters["timestamp"]),
			},
		},
	}

	result, getErr := svc.Query(queryInput)
	if getErr != nil {
		panic(getErr)
	}

	jsonData, _ := json.Marshal(result.Items)

	resp := Response{
		StatusCode:       200,
		IsBase64Encoded: false,
		Body:            string(jsonData),
		Headers: map[string]string{
			"Content-type":           "application/json",
			"Access-Control-Allow-Origin": "*",
		},
	}
	return resp, nil
}

func main() {
	lambda.Start(Handler)
}

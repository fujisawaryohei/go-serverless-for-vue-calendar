package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/fujisawaryohei/go-serverless-for-vue-calendar/todo"
)

type Request events.APIGatewayProxyRequest
type Response events.APIGatewayProxyResponse

func Handler(ctx context.Context, request Request) (Response, error) {
	svc := todo.NewSession()
	queryInput := &dynamodb.QueryInput{
		TableName: aws.String("my-vue-calendar-db"),
		ExpressionAttributeNames: map[string]*string{
			"#timestamp": aws.String("timestamp"),
		},
		KeyConditionExpression: aws.String("#timestamp = :timestamp"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":timestamp": {
				S: aws.String(request.QueryStringParameters["timestamp"]),
			},
		},
	}

	result, getErr := svc.Query(queryInput)
	if getErr != nil {
		panic(getErr)
	}

	body := []todo.Item{}
	if err := dynamodbattribute.UnmarshalListOfMaps(result.Items, &body); err != nil {
		panic(fmt.Sprintf("failed to unmarshal Dynamodb Scan Items, %v", err))
	}

	jsonData, _ := json.Marshal(body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            string(jsonData),
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
		},
	}
	return resp, nil
}

func main() {
	lambda.Start(Handler)
}

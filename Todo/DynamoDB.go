package todo

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func NewSession() *dynamodb.DynamoDB {
	mySession := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1")},
	))
	return dynamodb.New(mySession)
}

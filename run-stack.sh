#!/bin/bash 
set -eu
aws cloudformation validate-template --template-body file://"$(pwd)"/cloudformation/dynamoDB.yml
aws cloudformation create-stack --stack-name dynamoDBStack --template-body file://"$(pwd)"/cloudformation/dynamoDB.yml
# aws cloudformation wait stack-create-complete --stack-name dynamoDBStack

package client

import (
	"fmt"

	"ESRS/user_server/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var (
	dynamoDBClient *dynamodb.DynamoDB
)

func InitDynamoDB() {
	cfg := config.GetConfig()
	if cfg == nil {
		panic(config.NilConfigError)
	}
	dynamoConfig := cfg.Dynamo
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(dynamoConfig.Region),
		Credentials: credentials.NewStaticCredentials(dynamoConfig.AccessID, dynamoConfig.AccessSecret, ""),
	})
	if err != nil {
		panic(fmt.Sprintf("Error creating session: %s", err))
	}

	// Create DynamoDB client
	dynamoDBClient = dynamodb.New(sess)
}

func GetDynamoDBClient() *dynamodb.DynamoDB {
	return dynamoDBClient
}

package dao

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
	"time"

	"ESRS/user_server/client"
	"ESRS/user_server/dao/entity"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

const UserTable = "ESRS_USER"

var userTableDAO *userTableDynamoDAO

// dao
type UserTableDao interface {
	GetByUID(ctx context.Context, userID string) (*entity.UserTable, error)
	Create(ctx context.Context, userName, email string) (string, error)
	GetByEmail(ctx context.Context, email string) (*entity.UserTable, error)
	Update(ctx context.Context, params UpdateUserParams) error
}

type UpdateUserParams struct {
	UID            string
	PreferredGenre []string
}

type userTableDynamoDAO struct {
	dbClient *dynamodb.DynamoDB
}

func GetUserTableDAO() UserTableDao {
	if userTableDAO == nil {
		dbClient := client.GetDynamoDBClient()
		userTableDAO = &userTableDynamoDAO{
			dbClient: dbClient,
		}
	}

	return userTableDAO
}

// Returns UserID
func (dao *userTableDynamoDAO) Create(ctx context.Context, userName string, email string) (string, error) {
	dbClient := client.GetDynamoDBClient()
	now := time.Now()
	nowUnix := now.Unix()
	item := entity.UserTable{
		UserID:     generateUserID(userName),
		UserName:   userName,
		Email:      email,
		CreateTime: nowUnix,
		UpdateTime: nowUnix,
	}

	marshaledItem, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		log.Printf("[userTableDynamoDAO-Create] Got error marshalling new movie item: %s", err)
		return "", err
	}

	_, putItemErr := dbClient.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(UserTable),
		Item:      marshaledItem,
	})
	if putItemErr != nil {
		log.Printf("[userTableDynamoDAO-Create] Got error calling PutItem: %s", err)
		return "", putItemErr
	}
	return item.UserID, nil
}

func (dao *userTableDynamoDAO) GetByUID(ctx context.Context, userID string) (*entity.UserTable, error) {
	dbClient := client.GetDynamoDBClient()
	filter := expression.Name("user_id").Equal(expression.Value(userID))

	expr, err := expression.NewBuilder().WithFilter(filter).Build()
	if err != nil {
		log.Printf("[userTableDynamoDAO-GetByUID] Got error building expression: %s", err)
		return nil, err
	}

	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(UserTable),
	}

	result, err := dbClient.Scan(params)
	if err != nil {
		log.Printf("[userTableDynamoDAO-GetByUID] Query API call failed: %s", err)
		return nil, err
	}

	for _, i := range result.Items {
		item := entity.UserTable{}
		err = dynamodbattribute.UnmarshalMap(i, &item)

		if err != nil {
			log.Printf("[userTableDynamoDAO-GetByUID] Got error unmarshalling: %s", err)
			return nil, err
		}
		return &item, nil
	}

	return nil, nil
}

func (dao *userTableDynamoDAO) GetByEmail(ctx context.Context, email string) (*entity.UserTable, error) {
	dbClient := client.GetDynamoDBClient()
	filter := expression.Name("email").Equal(expression.Value(email))

	expr, err := expression.NewBuilder().WithFilter(filter).Build()
	if err != nil {
		log.Printf("[userTableDynamoDAO-GetByEmail] Got error building expression: %s", err)
		return nil, err
	}

	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(UserTable),
	}

	result, err := dbClient.Scan(params)
	if err != nil {
		log.Printf("[userTableDynamoDAO-GetByEmail] Query API call failed: %s", err)
		return nil, err
	}

	for _, i := range result.Items {
		item := entity.UserTable{}
		err = dynamodbattribute.UnmarshalMap(i, &item)

		if err != nil {
			log.Printf("[userTableDynamoDAO-GetByEmail] Got error unmarshalling: %s", err)
			return nil, err
		}
		return &item, nil
	}

	return nil, nil
}

func (dao *userTableDynamoDAO) Update(ctx context.Context, params UpdateUserParams) error {
	dbClient := client.GetDynamoDBClient()
	marshaledGenreList, err := dynamodbattribute.MarshalList(params.PreferredGenre)
	if err != nil {
		log.Printf("[userTableDynamoDAO-Update] Got error marshalling preferred genres: %s", err)
		return err
	}

	nowUnix := time.Now().Unix()
	// Construct the update input
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(UserTable),
		Key: map[string]*dynamodb.AttributeValue{
			"user_id": {
				S: aws.String(params.UID),
			},
		},
		UpdateExpression: aws.String("set preferred_genre = :g, update_time = :t"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":g": {
				L: marshaledGenreList,
			},
			":t": {
				N: aws.String(strconv.FormatInt(nowUnix, 10)),
			},
		},
		ReturnValues: aws.String("UPDATED_NEW"),
	}

	_, err = dbClient.UpdateItem(input)
	if err != nil {
		log.Printf("[userTableDynamoDAO-Update] Got error updating item: %s", err)
		return err
	}

	return nil
}

func generateUserID(username string) string {
	currentTime := time.Now().UnixNano()
	uniqueString := fmt.Sprintf("%s%d", username, currentTime)

	// Create a new MD5 hash
	hasher := md5.New()
	hasher.Write([]byte(uniqueString))

	// Get the MD5 hash as a hexadecimal string
	userID := hex.EncodeToString(hasher.Sum(nil))

	return userID
}

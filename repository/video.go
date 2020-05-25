package repository

import (
	"CourseX/entity"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type dynamoDBRepo struct {
	tableName string
}

func NewDynamoDBRepository() VideoRepository {
	return &dynamoDBRepo{
		tableName: "videos",
	}
}

func createDynamoDBClient() *dynamodb.DynamoDB {
	// Create AWS Session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Return DynamoDB client
	return dynamodb.New(sess)
}

func (repo *dynamoDBRepo) Save(video *entity.Video) (*entity.Video, error) {
	// Get a new DynamoDB client
	dynamoDBClient := createDynamoDBClient()

	// Transforms the Video to map[string]*dynamodb.AttributeValue
	attributeValue, err := dynamodbattribute.MarshalMap(video)
	if err != nil {
		return nil, err
	}

	// Create the Item Input
	item := &dynamodb.PutItemInput{
		Item:      attributeValue,
		TableName: aws.String(repo.tableName),
	}

	// Save the Item into DynamoDB
	_, err = dynamoDBClient.PutItem(item)
	if err != nil {
		return nil, err
	}

	return video, err
}

func (repo *dynamoDBRepo) FindAll() ([]entity.Video, error) {
	// Get a new DynamoDB client
	dynamoDBClient := createDynamoDBClient()

	// Build the query input parameters
	params := &dynamodb.ScanInput{
		TableName: aws.String(repo.tableName),
	}

	// Make the DynamoDB Query API call
	result, err := dynamoDBClient.Scan(params)
	if err != nil {
		return nil, err
	}
	var videos []entity.Video = []entity.Video{}
	for _, i := range result.Items {
		video := entity.Video{}

		err = dynamodbattribute.UnmarshalMap(i, &video)

		if err != nil {
			panic(err)
		}
		videos = append(videos, video)
	}
	return videos, nil
}

func (repo *dynamoDBRepo) FindByID(id string) (*entity.Video, error) {
	// Get a new DynamoDB client
	dynamoDBClient := createDynamoDBClient()

	result, err := dynamoDBClient.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(repo.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				N: aws.String(id),
			},
		},
	})
	if err != nil {
		return nil, err
	}
	video := entity.Video{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &video)
	if err != nil {
		panic(err)
	}
	return &video, nil
}

// Delete: TODO
func (repo *dynamoDBRepo) Delete(video *entity.Video) error {
	return nil
}

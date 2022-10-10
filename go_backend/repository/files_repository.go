package repository

import (
	"sample_app/model"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type FilesRepository interface {
	FindAll() ([]model.FileEntity, error)
	Find(name string) (*model.FileEntity, error)
	Save(file *model.FileEntity) error
}

type DynamoFilesRepository struct {
	client *dynamodb.DynamoDB
	table  string
}

func NewDynamoFilesRepository(sess *session.Session, tableName string) *DynamoFilesRepository {
	dynamo := dynamodb.New(sess)
	return &DynamoFilesRepository{client: dynamo, table: tableName}
}

func (repo *DynamoFilesRepository) FindAll() (entities []model.FileEntity, err error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(repo.table),
	}

	scanOut, err := repo.client.Scan(input)
	if err != nil {
		return nil, err
	}
	err = dynamodbattribute.UnmarshalListOfMaps(scanOut.Items, &entities)
	if err != nil {
		return nil, err
	}
	return entities, nil
}

func (repo *DynamoFilesRepository) Find(name string) (file *model.FileEntity, err error) {
	file = &model.FileEntity{}
	result, err := repo.client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(repo.table),
		Key: map[string]*dynamodb.AttributeValue{
			"filename": {
				S: aws.String(name),
			},
		},
	})
	if err != nil {
		return nil, err
	}
	err = dynamodbattribute.UnmarshalMap(result.Item, file)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (repo *DynamoFilesRepository) Save(file *model.FileEntity) error {
	tm, err := dynamodbattribute.MarshalMap(file)
	if err != nil {
		return err
	}
	input := &dynamodb.PutItemInput{
		Item:                tm,
		TableName:           aws.String(repo.table),
		ConditionExpression: aws.String("attribute_not_exists(filename)"),
	}
	_, err = repo.client.PutItem(input)
	if err != nil {
		return err
	}
	return nil
}

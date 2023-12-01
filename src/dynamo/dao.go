package dynamo

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

//region Errors
var ErrorNoItemReturned = errors.New("no item returned")
//endregion

type DynamoDao struct {
	DynamoApi DynamoApi
	TableName string
}

type DynamoApi interface {
	GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)
}

//region Constructors
func NewDynamoDaoFromConfig(cfg aws.Config, tableName string) DynamoDao {
	return DynamoDao{
		DynamoApi: dynamodb.NewFromConfig(cfg),
		TableName: tableName,
	}
}
//endregion

// GetById
func (d *DynamoDao) GetById(ctx context.Context, id string) (map[string]interface{}, error) {
	input := &dynamodb.GetItemInput{
		Key: map[string]types.AttributeValue{
			"Id": &types.AttributeValueMemberS{ Value: id},
			"Version": &types.AttributeValueMemberS{ Value: "1"},
		},
		TableName: &d.TableName,
	}

	resp, err := d.DynamoApi.GetItem(ctx, input)
	if err != nil {
		return nil, err
	}

	if resp.Item == nil || len(resp.Item) == 0 {
		return nil, ErrorNoItemReturned
	}

	var respValue map[string]interface{} 
	err = attributevalue.UnmarshalMap(resp.Item, &respValue)
	if err != nil {
		return nil, err
	}
	return respValue, nil
}
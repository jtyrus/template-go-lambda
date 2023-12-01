package dynamo

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/stretchr/testify/assert"
)

//region Mocks
type MockDynamoApi struct{}
var getItemOutput *dynamodb.GetItemOutput
var getItemOutputError error

func (m *MockDynamoApi) GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error) {
	getItemInputSpy = params

	return getItemOutput, getItemOutputError
}

//endregion

//region Argument Captors
var getItemInputSpy *dynamodb.GetItemInput
//endregion

var dao DynamoDao = DynamoDao{ DynamoApi: &MockDynamoApi{}, TableName: "test-table" }

func Test_GetById(t *testing.T) {
	setUp()
	givenGetItemReturns(validItem(), nil)

	resp, err := dao.GetById(context.Background(), "test-id")

	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "test-id", resp["Id"])
	assert.Equal(t, "test-table", *getItemInputSpy.TableName)
}

func Test_GetById_NilItem(t *testing.T) {
	setUp()
	givenGetItemReturns(nilItem(), nil)

	resp, err := dao.GetById(context.Background(), "test-id")

	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.Equal(t, "no item returned", err.Error())
	assert.Equal(t, "test-table", *getItemInputSpy.TableName)
}

func Test_GetById_EmptyResponse(t *testing.T) {
	setUp()
	givenGetItemReturns(missingItem(), nil)

	resp, err := dao.GetById(context.Background(), "test-id")

	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.Equal(t, "no item returned", err.Error())
	assert.Equal(t, "test-table", *getItemInputSpy.TableName)
}

func setUp() {
	getItemInputSpy = nil
	getItemOutput = nil
	getItemOutputError = nil
}

func givenGetItemReturns(resp *dynamodb.GetItemOutput, err error) {
	getItemOutput = resp
	getItemOutputError = err
}

//region responses
func validItem() *dynamodb.GetItemOutput {
	itemVal := map[string]interface{}{
		"Id": "test-id",
		"data": map[string]interface{}{
			"title":       "Cloud DVR",
			"subtitle":    "DirecTV Now (DirecTV Stream)",
			"description": "Lorem ipsum dolor sit amet consectetur adipisicing elit. Praesentium dolore rerum laborum iure enim sint nemo omnis voluptate exercitationem eius?",
			"image":       "https://upload.wikimedia.org/wikipedia/en/thumb/1/15/DirecTV_Now.svg/418px-DirecTV_Now.svg.png?20180527193140",
			"link":        "https://streamtv.directv.com/",
		},
	}

	item, _ := attributevalue.MarshalMap(itemVal)
	return &dynamodb.GetItemOutput{
		Item: item,
	}
}

func nilItem() *dynamodb.GetItemOutput {
	return &dynamodb.GetItemOutput{}
}

func missingItem() *dynamodb.GetItemOutput {
	item := map[string]types.AttributeValue{}
	return &dynamodb.GetItemOutput{ Item: item}
}
//endregion
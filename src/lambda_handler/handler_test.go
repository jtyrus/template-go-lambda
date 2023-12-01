package main

import (
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

//region Mocks
type MockDao struct {}

func (mDao *MockDao) GetById(ctx context.Context, id string) (map[string]interface{}, error) {
	getByIdInput = &id
	return mockDaoResponse, mockDaoError
}
//endregion

var getByIdInput *string
var mockDaoResponse map[string]interface{}
var mockDaoError error

func Test_GetData(t *testing.T) {
	setUp()
	input := events.APIGatewayProxyRequest{
		QueryStringParameters: map[string]string { "id": "test-id"},
	}

	mockDaoResponse = map[string]interface{} {
		"projects": []map[string]interface{}{ {"test-field": "test-value"} } ,
	}
	response, _ := getData(context.Background(), input)

	assert.Equal(t, "test-id", *getByIdInput)
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "{\"projects\":[{\"test-field\":\"test-value\"}]}", response.Body)
	assert.Equal(t, "*", response.Headers["Access-Control-Allow-Origin"])
}

func Test_GetData_MissingId(t *testing.T) {
	setUp()
	input := events.APIGatewayProxyRequest{
		QueryStringParameters: map[string]string { "id-field-2": "test-id"},
	}

	mockDaoResponse = map[string]interface{} {
		"test-field": "test-value",
	}
	response, _ := getData(context.Background(), input)

	assert.Nil(t, getByIdInput)
	assert.Equal(t, 400, response.StatusCode)
}

func setUp() {
	mockDaoResponse = nil
	mockDaoError = nil
	getByIdInput = nil

	deps = Dependencies{
		Dao: &MockDao{},
	}
}
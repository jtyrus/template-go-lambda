package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/jtyrus/portfolio_api/dynamo"
)

type Dependencies struct {
	Dao Dao
}

type Dao interface {
	GetById(ctx context.Context, id string) (map[string]interface{}, error)
}

var deps Dependencies

func main() {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Fatalf("failed to load configuration, %v", err)
	}

	dao := dynamo.NewDynamoDaoFromConfig(cfg, os.Getenv("TABLE_NAME"))

	deps = Dependencies{
		Dao: &dao,
	}
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(getData)
}

func getData(ctx context.Context,input events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println(input)
	response := events.APIGatewayProxyResponse { StatusCode: 500, Headers: map[string]string { "Access-Control-Allow-Origin": "*"} }
	id := input.QueryStringParameters["id"]

	if id == "" {
		response.StatusCode = 400
		return response, nil
	}

	if daoResp, daoErr := deps.Dao.GetById(ctx, id); daoErr != nil {
		response.Body = fmt.Sprint(daoErr)
		response.StatusCode = 500
	} else {
		jsonString, _ := json.Marshal(daoResp)
		response.Body = string(jsonString)
		response.StatusCode = 200
	}
	
	return response, nil
}

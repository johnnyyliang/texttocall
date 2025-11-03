package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	log.Println("Request received for makecall!")
	log.Printf("Received body: %s", request.Body)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Request successfully logged.",
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}

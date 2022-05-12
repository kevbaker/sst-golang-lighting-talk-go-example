package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	message := "Hello, World! Your request was received at " + request.RequestContext.Time + "!!"
	return events.APIGatewayProxyResponse{
		Body:       message,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}

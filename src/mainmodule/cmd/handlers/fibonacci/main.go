package main

import (
	"encoding/json"
	"net/http"

	"mainmodule/internal/services/fibonacci"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {

	s := fibonacci.FibonacciService{}
	fibonacciResponse := s.GetSequence()

	// Format results and return as API Gateway Response
	responseJSON, _ := json.Marshal(fibonacciResponse)

	// Build Response
	response := events.APIGatewayProxyResponse{
		Body:       string(responseJSON),
		StatusCode: http.StatusOK,
		Headers:    make(map[string]string),
	}
	response.Headers["Content-Type"] = "application/json"

	return response, nil

}

func main() {
	lambda.Start(Handler)
}

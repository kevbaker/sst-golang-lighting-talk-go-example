package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"mainmodule/internal/services/fibonacci"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	// Start time
	start := time.Now()

	// Settings
	DEFAULT_COUNT := 5

	// Get param for count
	count, err := strconv.Atoi(request.PathParameters["count"])
	if err != nil {
		count = DEFAULT_COUNT
	}

	// Get sequence
	s := fibonacci.FibonacciService{}
	fibonacciResponse := s.GetSequence(count)

	// Duration
	duration := time.Since(start)
	fibonacciResponse.Time = duration.Microseconds()

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

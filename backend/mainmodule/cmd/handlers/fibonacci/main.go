package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"mainmodule/internal/controllers"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {

	// Settings
	const DEFAULT_COUNT int = 5
	const DEFAULT_SAMPLERATE int = 1000

	// Get param for count
	count, err := strconv.Atoi(request.PathParameters["count"])
	if err != nil {
		count = DEFAULT_COUNT
	}

	// Get param for samplerate
	samplerate, err := strconv.Atoi(request.PathParameters["samplerate"])
	if err != nil {
		samplerate = DEFAULT_SAMPLERATE
	}

	// Get sequence
	c := controllers.DataController{}
	dataResponse := c.GetData(count, samplerate)

	// Format results and return as API Gateway Response
	responseJSON, _ := json.Marshal(dataResponse)

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

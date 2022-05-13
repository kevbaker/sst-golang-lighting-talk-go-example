package controllers

import (
	"mainmodule/internal/services/fibonacci"
	"mainmodule/internal/services/plotdata"
)

type DataResponse struct {
	Name      string                      `json:"name"`
	Fibonacci fibonacci.FibonacciResponse `json:"fibonacci"`
	Plotdata  plotdata.PlotdataResponse   `json:"plotdata"`
}

// DataController struct to contains features
// related to FibonacciService data
// NOTE: A controller can aggregate data from multiple services
// to be formatted by the cmd/transport layer
type DataController struct {
	Name string
}

// GetSequence of fibonacci numbers in a struct
// return a FibonacciResponse
func (c DataController) GetData(count int, sampleRate int) DataResponse {

	// Get sequence
	s := fibonacci.FibonacciService{}
	fibonacciResponse := s.GetSequence(count)

	// Get sequence
	ps := plotdata.PlotdataService{}
	plotdataResponse := ps.GetData(sampleRate)

	// Combined data
	response := DataResponse{
		Name:      "sst-golang-lighting-talk-go-example",
		Fibonacci: fibonacciResponse,
		Plotdata:  plotdataResponse}
	return response
}

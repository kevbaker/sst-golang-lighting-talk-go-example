package controllers

import (
	"mainmodule/internal/services/fibonacci"
)

// FibonacciController struct to contains features
// related to FibonacciService data
// NOTE: A controller can aggregate data from multiple services
// to be formatted by the cmd/transport layer
type FibonacciController struct {
	Name string
}

// GetSequence of fibonacci numbers in a struct
// return a FibonacciResponse
func (c FibonacciController) GetSequence(count int) fibonacci.FibonacciResponse {

	// Get sequence
	s := fibonacci.FibonacciService{}
	fibonacciResponse := s.GetSequence(count)

	return fibonacciResponse
}

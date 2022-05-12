package fibonacci

import "time"

// FibonacciService struct to contains features
// related to FibonacciService data
type FibonacciService struct {
	Name string
}

func (s FibonacciService) DoTest() []string {
	list := []string{"hello", "fibonacci"}
	return list
}

func (s FibonacciService) GetSequence() FibonacciResponse {
	response := FibonacciResponse{Time: time.Now().String(), Error: "NOPE", Max: 12, Sequence: []int{2, 3, 5, 7, 11, 13}}
	return response
}

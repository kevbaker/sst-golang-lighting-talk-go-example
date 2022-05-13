package fibonacci

// FibonacciService struct to contains features
// related to FibonacciService data
type FibonacciService struct {
	Name string
}

// GetSequence of fibonacci numbers in a struct
// return a FibonacciResponse
func (s FibonacciService) GetSequence(count int) FibonacciResponse {
	fibonacciSequence := s.getSequenceValues(count)
	response := FibonacciResponse{
		Name:     "fibonacci",
		Error:    "",
		Max:      fibonacciSequence[len(fibonacciSequence)-1],
		Sequence: fibonacciSequence}

	return response
}

// getSequenceValues of fibonacci numbers
// return a int array
func (s FibonacciService) getSequenceValues(count int) []int {
	fibonacci := []int{0, 1}
	for i := 1; i < count; i++ {
		next := (fibonacci[i] + fibonacci[i-1])
		fibonacci = append(fibonacci, next)
	}

	return fibonacci
}

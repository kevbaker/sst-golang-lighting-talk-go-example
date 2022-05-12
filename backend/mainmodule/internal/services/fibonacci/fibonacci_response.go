package  fibonacci

type FibonacciResponse struct {
	Time     string `json:"time"`
	Error    string `json:"error,omitempty"`
	Max      int    `json:"max"`
	Sequence []int  `json:"sequence"`
}

package fibonacci

type FibonacciResponse struct {
	Name     string `json:"name"`
	Time     string `json:"time"`
	Error    string `json:"error,omitempty"`
	Max      int    `json:"max"`
	Sequence []int  `json:"sequence"`
}

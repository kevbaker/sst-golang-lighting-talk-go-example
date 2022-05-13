package fibonacci

type FibonacciResponse struct {
	Name     string  `json:"name"`
	Time     float64 `json:"time"`
	Error    string  `json:"error,omitempty"`
	Max      int     `json:"max"`
	Sequence []int   `json:"sequence"`
}

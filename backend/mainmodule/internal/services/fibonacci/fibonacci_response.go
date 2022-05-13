package fibonacci

type FibonacciResponse struct {
	Name     string `json:"name"`
	Time     int64  `json:"time"`
	Error    string `json:"error,omitempty"`
	Max      int    `json:"max"`
	Sequence []int  `json:"sequence"`
}

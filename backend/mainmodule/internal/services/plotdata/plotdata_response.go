package plotdata

type PlotdataResponse struct {
	Name  string    `json:"name,omitempty"`
	Time  float64   `json:"time,omitempty"`
	Error string    `json:"error,omitempty"`
	Data  []float64 `json:"data"`
}

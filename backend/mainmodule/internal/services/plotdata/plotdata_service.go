package plotdata

import (
	"fmt"
	"math"
)

// PlotdataService struct to contains features
// related to PlotdataService data
type PlotdataService struct {
	Name string
}

// pDuration           = 5
// pSampleRate         = 44000
// pFrequency  float64 = 100

// GetData for plotdata numbers in a struct
// return a PlotdataResponse
func (s PlotdataService) GetData(samplerate int) PlotdataResponse {

	response := PlotdataResponse{Name: "plotdata"}
	response.Data = s.generate(5, samplerate, 100)

	return response
}

func (s PlotdataService) generate(Duration int, SampleRate int, Frequency float64) []float64 {
	var (
		start float64 = 1.0
		end   float64 = 1.0e-4
		out   []float64
	)
	nsamps := Duration * SampleRate
	tau := math.Pi * 2
	var angle float64 = tau / float64(nsamps)
	decayfac := math.Pow(end/start, 1.0/float64(nsamps))
	for i := 0; i < nsamps; i++ {
		sample := math.Sin(angle * Frequency * float64(i))
		sample *= start
		start *= decayfac
		fmt.Printf("%.8f\n", sample)
		out = append(out, sample)
	}

	return out
}

package main

import (
	"fmt"
	"math"

	"github.com/sbinet/go-gnuplot"
)

// REF: https://github.com/sbinet/go-gnuplot

const (
	pDuration           = 5
	pSampleRate         = 44000
	pFrequency  float64 = 100
)

func generate(Duration int, SampleRate int, Frequency float64) []float64 {
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

func main() {
	fname := ""
	persist := false
	debug := true

	p, err := gnuplot.NewPlotter(fname, persist, debug)
	if err != nil {
		err_string := fmt.Sprintf("** err: %v\n", err)
		panic(err_string)
	}
	defer p.Close()

	values := generate(pDuration, pSampleRate, pFrequency)
	p.PlotX(values, "amazing data")
	p.CheckedCmd("set terminal png")
	p.CheckedCmd("set output 'plot.png'")
	p.CheckedCmd("replot")

	p.CheckedCmd("q")

}

package main

// gnuplot cli application
// This is a simple app to demonstrate usign the same controller code
// from the AWS Lambda function in a CLI

import (
	"fmt"

	"mainmodule/internal/controllers"

	"github.com/sbinet/go-gnuplot"
)

// REF: https://github.com/sbinet/go-gnuplot

func main() {

	// Call controller for data
	c := controllers.DataController{}
	dataResponse := c.GetData(100, 1000)
	values := dataResponse.Plotdata.Data

	// Initialize Plotter
	fname := ""
	persist := false
	debug := true
	p, err := gnuplot.NewPlotter(fname, persist, debug)
	if err != nil {
		err_string := fmt.Sprintf("** err: %v\n", err)
		panic(err_string)
	}
	defer p.Close()

	// Send Values to Plot
	p.PlotX(values, "amazing data")

	// Define format and run plot
	p.CheckedCmd("set terminal png")
	p.CheckedCmd("set output 'plot.png'")
	p.CheckedCmd("replot")
	p.CheckedCmd("q")

}

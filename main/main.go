package main

import (
	"aisd-grafy/data"
	"aisd-grafy/euler"
	"aisd-grafy/hamilton"
)

type Result struct {
	Time float64
	Name string
}

func main() {

	var graph [][]int = data.GetData()

	euler.Euler(0, &graph)
	hamilton.Hamilton(0, &graph)

	// var timeValue float64
	// start := time.Now()
	// elapsed := time.Since(start)

}
